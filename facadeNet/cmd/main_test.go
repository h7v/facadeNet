package main

// Tests for main.go

import (
	"testing"

	"github.com/h7v/facadeNet/pkg/facade"
	"github.com/h7v/facadeNet/pkg/set-server"
	"github.com/h7v/facadeNet/pkg/work-server"
)

func TestOk(t *testing.T) {
	set := set_server.NewSetServer()
	okResult := "set server"
	result := set.CheckSetServer()
	if result != okResult {
		t.Errorf("Test for valid argument failed!\nExpected:\n%v\nGot:\n%v", okResult, result)
	}
	work := work_server.NewWorkServer()
	okResult = "work server"
	result = work.CheckWorkServer()
	if result != okResult {
		t.Errorf("Test for valid argument failed!\nExpected:\n%v\nGot:\n%v", okResult, result)
	}
	f := facade.WorkingServer(set, work)
	f.Work()
}

// TestSeparatedMockTestForSetServer is separated mock test for SetServer
func TestSeparatedMockTestForSetServer(t *testing.T) {
	set := set_server.NewSetServer()
	okResult := "set server"
	result := set.CheckSetServer()
	if result != okResult {
		t.Errorf("Test for valid argument failed!\nExpected:\n%v\nGot:\n%v", okResult, result)
	}
}

// TestSeparatedMockTestForWorkServer is separated mock test for WorkServer
func TestSeparatedMockTestForWorkServer(t *testing.T) {
	work := work_server.NewWorkServer()
	okResult := "work server"
	result := work.CheckWorkServer()
	if result != okResult {
		t.Errorf("Test for valid argument failed!\nExpected:\n%v\nGot:\n%v", okResult, result)
	}
}

