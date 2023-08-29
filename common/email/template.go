package email

const (
	SendCodeTemplate = `
		<!DOCTYPE html>
        <html>
        <head>
            <style>
                body {
                    font-family: Arial, sans-serif;
                    background-color: #f5f5f5;
                }
                .container {
                    max-width: 600px;
                    margin: 0 auto;
                    padding: 20px;
                    background-color: #ffffff;
                    border-radius: 10px;
                    box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);
                }
                .title {
                    font-size: 24px;
                    color: #333333;
                    margin-bottom: 20px;
                }
                .code {
                    font-size: 32px;
                    color: #007bff;
                }
                .btn-container {
                    margin-top: 20px;
                }
                .btn {
                    display: inline-block;
                    padding: 10px 20px;
                    background-color: #dddddd;
                    color: #333333;
                    text-decoration: none;
                    border-radius: 5px;
                    transition: background-color 0.3s;
                }
                .btn:hover {
                    background-color: #cccccc;
                }
            </style>
        </head>
        <body>
            <div class="container">
                <div class="title">OurAI 验证码</div>
                <p>以下是您本次请求的验证码，验证码有效期为10分钟，请勿透露给他人:</p>
                <div class="code">%s</div>
                <div class="btn-container">
                    <a class="btn" href="https://www.ourai.club">返回OurAI官网</a>
                </div>
            </div>
        </body>
        </html>
`
)
