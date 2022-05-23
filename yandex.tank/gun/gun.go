// create a package
package main

// import some pandora stuff
// and stuff you need for your scenario
// and protobuf contracts for your grpc service

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	pb "main/my_contracts"

	uuid "github.com/google/uuid"
	"github.com/spf13/afero"
	"github.com/yandex/pandora/cli"
	phttp "github.com/yandex/pandora/components/phttp/import"
	"github.com/yandex/pandora/core"
	"github.com/yandex/pandora/core/aggregator/netsample"
	coreimport "github.com/yandex/pandora/core/import"
	"github.com/yandex/pandora/core/register"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/durationpb"
)

type Ammo struct {
	VideoId       string
	UserId        string
	EventType     string
	EventDuration int
	EventTime     string
}

type Sample struct {
	URL              string
	ShootTimeSeconds float64
}

type GunConfig struct {
	Target string `validate:"required"`
}

type Gun struct {
	// Configured on construction.
	client grpc.ClientConn
	conf   GunConfig
	// Configured on Bind, before shooting
	aggr core.Aggregator // May be your custom Aggregator.
	core.GunDeps
}

func NewGun(conf GunConfig) *Gun {
	return &Gun{conf: conf}
}

func (g *Gun) Bind(aggr core.Aggregator, deps core.GunDeps) error {
	conn, err := grpc.Dial(
		g.conf.Target,
		grpc.WithInsecure(),
		grpc.WithTimeout(time.Second),
		grpc.WithUserAgent("load test, pandora custom shooter"))
	if err != nil {
		log.Fatalf("FATAL: %s", err)
	}
	g.client = *conn
	g.aggr = aggr
	g.GunDeps = deps
	return nil
}

func (g *Gun) Shoot(ammo core.Ammo) {
	customAmmo := ammo.(*Ammo)
	g.shoot(customAmmo)
}

func (g *Gun) case1_method(client pb.KionServiceClient, ammo *Ammo) int {
	events := [4]string{"STOP_VIDEO_EVENT", "FORWARD_VIDEO_EVENT", "REVERT_VIDEO_EVENT", "EXIT_VIDEO_EVENT"}
	rand.Seed(time.Now().Unix())
	duration := durationpb.Duration{Seconds: int64(rand.Intn(100)), Nanos: 0}
	response, err := client.CreateRecord(context.Background(), &pb.CreateRecordRequest{
		VideoId:   uuid.UUID(uuid.New()).String(),
		UserId:    uuid.UUID(uuid.New()).String(),
		EventType: events[rand.Intn(len(events))],
		Time:      &duration})
	fmt.Println(response)
	if err != nil {
		fmt.Println(err)
		return 500
	}
	return 200
}

func (g *Gun) case2_method(client pb.KionServiceClient, ammo *Ammo) int {
	response, err := client.GetLatestRecord(context.Background(), &pb.GetLatestRecordRequest{
		UserId:  ammo.UserId,
		VideoId: ammo.VideoId,
	})

	fmt.Println(response)
	if err != nil {
		fmt.Println(err)
		return 500
	}
	return 200
}

func (g *Gun) shoot(ammo *Ammo) {
	code := 0
	sample := netsample.Acquire(ammo.VideoId)

	conn := g.client
	client := pb.NewKionServiceClient(&conn)

	// Writing
	code = g.case1_method(client, ammo)

	// Reading
	// code = g.case2_method(client, ammo)

	defer func() {
		sample.SetProtoCode(code)
		g.aggr.Report(sample)
	}()
}

func main() {
	fs := afero.NewOsFs()
	coreimport.Import(fs)

	phttp.Import(fs)

	coreimport.RegisterCustomJSONProvider("custom_provider", func() core.Ammo { return &Ammo{} })

	register.Gun("My_custom_gun_name", NewGun, func() GunConfig {
		return GunConfig{
			Target: "localhost:8080",
		}
	})

	cli.Run()
}
