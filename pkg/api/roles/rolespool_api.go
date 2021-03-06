package roles

import (
	"encoding/json"
	"errors"
	"users/pkg/api"
	"users/pkg/domain/roles"
)

const rolesurl = api.BASE_URL + "/roles/pool"

type RolesPool roles.RolesPool

func getRolespools() ([]RolesPool, error) {
	entities := []RolesPool{}
	resp, _ := api.Rest().Get(rolesurl + "/all")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}

func getRolespool(id string) (RolesPool, error) {
	rolespool := RolesPool{}
	resp, _ := api.Rest().Get(rolesurl + "/get/" + id)
	if resp.IsError() {
		return rolespool, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &rolespool)
	if err != nil {
		return rolespool, errors.New(resp.Status())
	}
	return rolespool, nil

}

func createRolespool(entity RolesPool) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(rolesurl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}
func updateRolespool(entity RolesPool) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(rolesurl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}

func deleteRolespool(entity RolesPool) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(rolesurl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}
