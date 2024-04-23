import requests as req
headers = {"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36"}
res = req.get("https://www.tutorialspoint.com/csharp/online-csharp-compiler.php", headers=headers)
print(res.text)