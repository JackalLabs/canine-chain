for i in {1..500}
do
  canined tx storage post-random 50 --home=$HOME/canine-test --from charlie --gas-prices=0.02ujkl -y
  sleep 20
done

