SUB_COMMAND="$1"
if [ $SUB_COMMAND = 'start' ]; then
  docker-compose up -d;
  #  timeout 10s bash -c 'until docker-compose exec db pg_isready; do sleep 2; done'
  # TODO: some issue with timeout line, runs fine on terminal so just waiting for 10 secs
  sleep 10
  exit
fi

if [ $SUB_COMMAND = 'stop' ]; then
  docker-compose down
  sleep 10
  exit
else
  echo "unknown command" $SUB_COMMAND
fi
