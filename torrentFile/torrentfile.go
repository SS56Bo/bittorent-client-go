package main

type TorrentFile struct {
	AnnounceString string
	InfoHash       [20]byte //	info hash is a cryptographic hash (usually SHA-1)
	PieceHash      [][20]byte
	PieceLength    int
	Length         int
	Name           string
}