package main

import structs "structs/photos"

type FBPhotos struct {
	Data   []structs.Data `json:"data"`
	Paging structs.Paging `json:"paging"`
}
