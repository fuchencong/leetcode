#bin/bash
cmd="insert into ipv6_allocations (ip_address) values "
i=0
j=0
for((i=0;i<16;i++));
do
for((j=0;j<16;j++));
do
ins_i=`printf %x $i`
ins_j=`printf %x $j`
cmd=${cmd}"(\"2403:ed40:e001::08${ins_i}${ins_j}\"),"
done
done
cmd="$(echo "$cmd" | rev | sed 's/,//' | rev)"
cmd=${cmd}";"
echo "$cmd"
