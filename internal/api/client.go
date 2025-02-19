package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/DavidinhoHD/agario/internal/game"
)
var url string = "localhost:8080"

// request all current food & player positions
func GetGameState() (game.GameState, error){
	var localGameState game.GameState
	resp, err := http.Get(url)
	if err != nil {
		return localGameState, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return localGameState, err
	}

	err = json.Unmarshal(b, &localGameState)
	if err != nil {
		return localGameState, err
	}

	return localGameState, nil
}

// request a new player from the server
func GetNewPlayer() (game.Player, error){
	var newPlayer game.Player
	resp, err := http.Get(url+"/init")
	if err != nil {
		return newPlayer, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return newPlayer, err
	}

	err  = json.Unmarshal(b, &newPlayer)
	if err != nil {
		return newPlayer, err
	}

	return newPlayer, nil
}

// update the player on the server
func UpdatePlayer(p game.Player) error{
	updatedPlayerJSON, err := json.Marshal(p)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", url+"/player", bytes.NewBuffer(updatedPlayerJSON))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// delete request to /food/{id}
func DeleteFoodById(f game.Food) error{
	addr := fmt.Sprintf("%v/food/%v", url, f.ID)
	req, err := http.NewRequest("DELETE", addr, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// delete request to /player/{id}
func DeletePlayerById(p game.Player) error{
	addr := fmt.Sprintf("%v/player/%v", url, p.ID)
	req, err := http.NewRequest("DELETE", addr, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}