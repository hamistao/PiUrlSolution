import requests

pi = requests.get("https://api.pi.delivery/v1/pi?start=0&numberOfDigits=1001&radix=10")

print(pi.content)