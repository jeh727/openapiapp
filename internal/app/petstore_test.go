package app_test

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jeh727/openapiapp/internal/app"
	"github.com/jeh727/openapiapp/internal/app/appapi"
	"gotest.tools/v3/golden"
)

func TestPetStoreFindPets(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		code   int
		params appapi.FindPetsParams
		golden string
	}{
		{name: "Happy Path", code: 200, params: appapi.FindPetsParams{}, golden: "find_pets1.golden"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequestWithContext(context.Background(), "GET", "/pets", nil)
			rec := httptest.NewRecorder()

			petStore := &app.PetStore{}
			petStore.FindPets(rec, req, testCase.params)

			code := rec.Code
			if code != testCase.code {
				t.Fatalf("expected status code %d, got %d", testCase.code, code)
			}

			res := rec.Body.String()
			golden.Assert(t, res, testCase.golden)
		})
	}
}

func TestPetStoreAddPet(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		req    string
		code   int
		golden string
	}{
		{
			name:   "Happy Path",
			code:   200,
			golden: "add_pet1.golden",
			req:    `{"id": 1, "name": "Fluffy", "tag": "cat"}`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequestWithContext(context.Background(), "POST", "/pets", strings.NewReader(testCase.req))
			rec := httptest.NewRecorder()

			petStore := &app.PetStore{}
			petStore.AddPet(rec, req)

			code := rec.Code
			if code != testCase.code {
				t.Fatalf("expected status code %d, got %d", testCase.code, code)
			}

			res := rec.Body.String()
			golden.Assert(t, res, testCase.golden)
		})
	}
}

func TestPetStoreDeletePet(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		req    string
		code   int
		golden string
	}{
		{
			name:   "Happy Path",
			code:   200,
			golden: "delete_pet1.golden",
			req:    `{"id": 1, "name": "Fluffy", "tag": "cat"}`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequestWithContext(context.Background(), "DELETE", "/pets/12345", strings.NewReader(testCase.req))
			rec := httptest.NewRecorder()

			petStore := &app.PetStore{}
			petStore.DeletePet(rec, req, 12345)

			code := rec.Code
			if code != testCase.code {
				t.Fatalf("expected status code %d, got %d", testCase.code, code)
			}

			res := rec.Body.String()
			golden.Assert(t, res, testCase.golden)
		})
	}
}

func TestPetStoreFindPetByID(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		code   int
		golden string
	}{
		{
			name:   "Happy Path",
			code:   200,
			golden: "find_pet_by_id1.golden",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequestWithContext(context.Background(), "GET", "/pets/12345", nil)
			rec := httptest.NewRecorder()

			petStore := &app.PetStore{}
			petStore.FindPetByID(rec, req, 12345)

			code := rec.Code
			if code != testCase.code {
				t.Fatalf("expected status code %d, got %d", testCase.code, code)
			}

			res := rec.Body.String()
			golden.Assert(t, res, testCase.golden)
		})
	}
}
