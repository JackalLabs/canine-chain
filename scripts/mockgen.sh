mockgen_cmd="mockgen"

$mockgen_cmd -source=x/rns/types/expected_keepers.go -package testutil -destination x/rns/testutil/expected_keepers_mocks.go
$mockgen_cmd -source=x/storage/types/expected_keepers.go -package testutil -destination x/storage/testutil/expected_keepers_mocks.go
