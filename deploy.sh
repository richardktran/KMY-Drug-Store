systemctl stop api.quaythuockhanhmy.service
cp api.quaythuockhanhmy.service /etc/systemd/system/api.quaythuockhanhmy.service
systemctl daemon-reload
systemctl start api.quaythuockhanhmy.service