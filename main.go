package main

import (
	"context"
	"errors"
	"log"

	"github.com/micro/go-micro/v2"
	pb "github.com/repodevs/shippy-service-vessel/proto/vessel"
)

// Repository ...
type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

// VesselRepository ...
type VesselRepository struct {
	vessels []*pb.Vessel
}

// FindAvailable - check a specifications against a map of vessels
// if capacity and max weight are bellow a vessels capacity and max weight,
// than return the vessel
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("No vessel found with that spec")
}

// Our gRPC service handler
type vesselService struct {
	repo Repository
}

// FindAvailable ...
func (s *vesselService) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}

	res.Vessel = vessel
	return nil
}

func main() {
	vessels := []*pb.Vessel{
		&pb.Vessel{Id: "vessel01", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	}

	repo := &VesselRepository{vessels}

	service := micro.NewService(
		micro.Name("shippy.service.vessel"),
	)
	service.Init()

	// Register implementation
	if err := pb.RegisterVesselServiceHandler(service.Server(), &vesselService{repo}); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
