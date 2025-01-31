package data

type Sender struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type Owner struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

var Admin Sender = Sender{
	Name:  "404th",
	Email: "umarov.doniyor.2003@gmail.com",
}

var Owners []Owner = []Owner{
	{Name: "Alex Kenzo", Email: "alex.coder.kenzo@gmail.com"},
	{Name: "Doniyor Umarov", Email: "umarovdoniyorstudy@gmail.com"},
	{Name: "Ismoil Safarov", Email: "umarov.doniyor.2001@gmail.com"},
}

var RandomEmailSubject string = `To'y qachon`

var RandomEmailBody string = `
	<body style="font-family: Arial, sans-serif; background-color: #f7f7f7; color: #333; margin: 0; padding: 0;">

  <div style="width: 100%; max-width: 600px; margin: 20px auto; background-color: #ffffff; border-radius: 8px; box-shadow: 0px 4px 12px rgba(0, 0, 0, 0.1); overflow: hidden;">

    <div style="background-color: #4CAF50; color: #ffffff; padding: 20px; text-align: center;">
      <h1 style="margin: 0; font-size: 24px;">Hello, Alex Kenzo!</h1>
    </div>

    <div style="padding: 20px;">
      <h2 style="color: #333333; font-size: 20px; margin-top: 0;">We’ve Got Exciting News Just for You!</h2>
      <p style="line-height: 1.6; font-size: 16px;">We hope this message finds you well! We’ve been working on some amazing updates to improve your experience, and we can’t wait to share them with you.</p>
      
      <h3 style="color: #4CAF50; font-size: 18px;">Here’s What’s New:</h3>
      <ul style="padding-left: 20px; font-size: 16px; line-height: 1.6;">
        <li><strong>Enhanced Features:</strong> Enjoy faster speeds, more customization, and even better performance!</li>
        <li><strong>Account Review:</strong> We recommend checking your account settings to get the most out of these changes.</li>
        <li><strong>Feedback Opportunity:</strong> We’d love to hear your thoughts to keep making improvements that matter to you.</li>
      </ul>

      <p style="line-height: 1.6; font-size: 16px;">If you have any questions, just reply to this email, and our team will be happy to help.</p>

      <div style="text-align: center; margin-top: 20px;">
        <a href="https://yourwebsite.com" style="background-color: #4CAF50; color: #ffffff; padding: 12px 20px; text-decoration: none; font-size: 16px; border-radius: 4px;">Explore Updates</a>
      </div>

      <p style="line-height: 1.6; font-size: 16px; color: #555555; text-align: center; margin-top: 20px;">
        Thank you for being part of our journey, Alex. We look forward to helping you make the most of our services!
      </p>
    </div>

    <div style="background-color: #f0f0f0; padding: 10px; text-align: center; color: #777777; font-size: 12px;">
      © 2024 Your Company, All rights reserved.
    </div>

  </div>
`
var RandomEmailIntro string = `
  <h3 style="color: #4CAF50; font-size: 18px;">Here’s a Quick Introduction</h3>
      <p style="line-height: 1.6; font-size: 16px;">We value our users and aim to provide a seamless experience. Below, you’ll find details on what’s new and how to make the most of our latest updates.</p>
`

type Mail struct {
	To                   string `json:"to"`
	ToName               string `json:"to_name"`
	From                 string `json:"from"`
	FromName             string `json:"from_name"`
	EmailBanner          string `json:"email_banner"`
	EmailIntroTopic      string `json:"email_intro_topic"`
	EmailIntroBody       string `json:"email_intro_body"`
	EmailBodyTopic       string `json:"email_body_topic"`
	EmailBody            string `json:"email_body"`
	EmailBodyConclusion  string `json:"email_body_conclusion"`
	EmailButtonName      string `json:"email_button_name"`
	EmailButtonUnderText string `json:"email_button_under_text"`
	EmailCopyright       string `json:"email_copyright"`
}

// --
// to whom
// to whom name
// from whom
// from whom name
// --
// --
// email banner
// --
// email intro theme
// email into
// --
// email body theme
// email body in <td>
// email body conclusion
// --
// button to website
// under button text part
// --
// ©️ copyright
// --
