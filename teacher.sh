id=$(grep 'Church' -r interviews -l | cut -d '-' -f2)
echo $id
cat interviews/interview-"$id"
echo $MAIN_SUSPECT