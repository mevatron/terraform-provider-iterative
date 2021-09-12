export PKR_VAR_test=true
export PKR_VAR_image_name=iterative-cml-$(cat /proc/sys/kernel/random/uuid)
packer build .
