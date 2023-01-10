package main

import "structs/albums"

type FBAlbums struct {
	Data []albums.Data `json:"data"`
	Paging albums.Paging `json:"paging"`
}
