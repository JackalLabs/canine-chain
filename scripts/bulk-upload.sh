for i in {1..75}
do
  canined tx storage post-random 50 --home=$HOME/canine-test --from charlie --gas-prices=0.02ujkl -y
  sleep 10
done

