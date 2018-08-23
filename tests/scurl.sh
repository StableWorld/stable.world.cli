
fail (){
  echo
  echo "[FAIL]" "$*"
  echo
}

test_that_bucket_envvar_is_required(){
  echo "test_that_bucket_envvar_is_required"
  {
    go run commands/scurl/main.go
  } && {
    fail "STABLE_WORLD_BUCKET should have been required"
  }
}

test_that_arguments_are_transformed(){
  echo "test_that_arguments_are_transformed"
  {
    STABLE_WORLD_BUCKET=12345 STABLE_WORLD_CURL=echo go run commands/scurl/main.go google.com
  } || {
    fail "STABLE_WORLD_BUCKET should have been required"
  }
}

main(){
  test_that_bucket_envvar_is_required
  test_that_arguments_are_transformed
}

main
