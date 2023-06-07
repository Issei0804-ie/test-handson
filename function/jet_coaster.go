package function

// この遊園地のジェットコースターには、年齢制限と身長制限があります。年齢制限は、12歳以上70歳未満です。身長制限は、130cm以上180cm未満です。
func CanRideJetCoaster(age int, height int) bool {
	return 12 <= age && age < 70 && 130 <= height && height < 180
}
