make ARGS=type/* all;
make ARGS=item/item all;
make ARGS=cart/cart all;

rm -rf ../pkg/rpc/*;
cp -r generated/github.com/sedeotech/backend-interview/pkg/rpc/* ../pkg/rpc/;