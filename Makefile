.PHONY:
prepare:
	echo "\n\n\n" >> /etc/hosts
	echo "192.168.0.202 mysql.test.com" >> /etc/hosts
	echo "192.168.0.202 redis.test.com" >> /etc/hosts
	echo "192.168.0.202 mongo.test.com" >> /etc/hosts
