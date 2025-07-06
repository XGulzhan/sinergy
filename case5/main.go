package main

import (
	"fmt"
	"os"
)

func main() {
	orgName := "АО \"Газпромбанк\""
	orgDescription := `АО "Газпромбанк" - один из крупнейших финансовых институтов России.
  За последние несколько лет был успешно реализован проект цифровой трансформации: перестроена IT-инфраструктура, 
  внедрены новые сервисы для улучшения клиентского опыта, перезапущены дистанционные каналы обслуживания.  
  Основан в 1990 году, штаб-квартира 
  расположена в Москве.`
	orgAddress := "117420, г. Москва, ул. Наметкина, дом 16, корпус 1."
	orgPhone := "+7 (495) 913-74-74"
	orgEmail := "mailbox@gazprombank.ru"
	orgSite := "https://www.gazprombank.ru"

	fonts := []struct {
		name    string
		family  string
		comment string
	}{
		{"Georgia", "'Georgia', serif", "Классический шрифт с засечками, хорошо подходит для официальных документов"},
		{"Roboto", "'Roboto', sans-serif", "Универсальный шрифт от Google, хорошая читаемость"},
		{"Open Sans", "'Open Sans', sans-serif", "Популярный шрифт с отличной разборчивостью"},
		{"Arial", "'Arial', sans-serif", "Стандартный шрифт, используется во многих компаниях"},
		{"Times New Roman", "'Times New Roman', serif", "Классический шрифт для официальных документов"},
	}
	fontLinks := `
  <link href="https://fonts.googleapis.com/css2?family=Roboto:wght@400;700&display=swap" rel="stylesheet">
  <link href="https://fonts.googleapis.com/css2?family=Open+Sans:wght@400;700&display=swap" rel="stylesheet">
  <link href="https://fonts.googleapis.com/css2?family=Times+New+Roman&display=swap" rel="stylesheet">`

	//   fontStyles := ""
	//   for i, font := range fonts {
	//     fontStyles += fmt.Sprintf(`
	// .font-option-%d {
	//   font-family: %s;
	// }`, i+1, font.family)
	//   }
	fontStyles := ""
	for i, font := range fonts {
		fontStyles += fmt.Sprintf(`
.font-option-%d {
  font-family: %s;
  font-size: 16px;
  line-height: 1.5;
  margin: 10px 0;
}`, i+1, font.family)
	}

	fontStyles += `
.font-option-5 {
  font-size: 15px;
  line-height: 1.4;
}`
	html := fmt.Sprintf(`<!DOCTYPE html>
<html lang="ru">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>%s</title>
  %s
  <style>
    body {
      font-family: 'Open Sans', sans-serif;
      line-height: 1.6;
      margin: 0;
      padding: 20px;
      color: #333;
      background-color: #f9f9f9;
    }
    .header {
      background-color: #0056b3;
      color: white;
      padding: 20px;
      text-align: center;
      margin-bottom: 30px;
      border-radius: 5px;
      box-shadow: 0 2px 5px rgba(0,0,0,0.1);
    }
    .org-info {
      background: white;
      padding: 25px;
      border-radius: 5px;
      box-shadow: 0 2px 5px rgba(0,0,0,0.1);
      margin-bottom: 30px;
    }
    .font-options {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
      gap: 20px;
      margin-top: 30px;
    }
    .font-option {
      padding: 15px;
      border: 1px solid #ddd;
      border-radius: 5px;
      background: white;
    }
    .font-label {
      font-weight: bold;
      margin-bottom: 10px;
      color: #0056b3;
    }
    .font-comment {
      font-size: 0.9em;
      color: #666;
      margin-top: 5px;
    }
    .info-item {
      margin-bottom: 10px;
    }
    .info-label {
      font-weight: bold;
    }
    %s
  </style>
</head>
<body>
  <div class="header">
    <h1>Наименование организации на базе, которой Вы проходите практическую подготовку</h1>
  </div>

  <div class="org-info">
    <h2>%s</h2>
    
    <div class="info-item">
      <span class="info-label">Адрес:</span> %s
    </div>
    <div class="info-item">
      <span class="info-label">Телефон:</span> %s
    </div>
    <div class="info-item">
      <span class="info-label">Email:</span> %s
    </div>
    <div class="info-item">
      <span class="info-label">Сайт:</span> <a href="%s"> %s </a></span>
    </div>
  </div>

<div class="org-info">
    <h2>Варианты шрифтов для организации</h2>
    <p>Ниже представлены 5 вариантов шрифтов, которые могут подойти для официального использования.</p>
    
    <div class="font-options">`, orgName, fontLinks, fontStyles, orgDescription, orgAddress, orgPhone, orgEmail, orgSite, orgSite)

	for i, font := range fonts {
		html += fmt.Sprintf(`
      <div class="font-option">
        <div class="font-label">Вариант %d: %s</div>
        <div class="font-option-%d">
          Пример текста на шрифте %s. Этот шрифт может использоваться для официальных документов компании.<br>
        </div>
        <div class="font-comment">%s</div>
      </div>`, i+1, font.name, i+1, font.name, font.comment)
	}

	html += `
    </div>
  </div>
</body>
</html>`

	err := os.WriteFile("organization_page.html", []byte(html), 0644)
	if err != nil {
		fmt.Println("Ошибка при сохранении файла:", err)
		return
	}

	fmt.Println("HTML-страница успешно создана: organization_page.html")
}
