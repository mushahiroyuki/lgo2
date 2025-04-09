echo "1. すぐに実行"
curl http://localhost:8080/request &
echo "2. すぐに実行"
curl http://localhost:8080/request &
echo "3. すぐに実行"
curl http://localhost:8080/request &
echo "4. すぐに実行"
curl http://localhost:8080/request &
echo "5. すぐに実行"
curl http://localhost:8080/request &
echo "6. すぐに実行"
curl http://localhost:8080/request &
echo "7. すぐに実行"
curl http://localhost:8080/request &
echo "8. すぐに実行"
curl http://localhost:8080/request &
echo "2秒sleepします"
sleep 2
echo "9. 2秒sleep後の実行"
curl http://localhost:8080/request
echo "3秒sleepします"
sleep 3
echo "10. 3秒sleep後の実行"
curl http://localhost:8080/request
