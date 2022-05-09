package bitlink

type BitLinkRepository interface {
	Find(hash string) *BitLink
	Save(bitlink *BitLink) bool
}
