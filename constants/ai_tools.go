package constants

import "github.com/sashabaranov/go-openai"

var Tools = []openai.Tool{
	{
		Type: openai.ToolTypeFunction,
		Function: &openai.FunctionDefinition{
			Name:        "get_menu_queries",
			Description: "Извлекает запросы к меню, которые нужно проверить в векторной базе",
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"items": map[string]interface{}{
						"type": "array",
						"items": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"query": map[string]interface{}{
									"type":        "string",
									"description": "Что именно ищет пользователь (например: пицца, хинкали, сок)",
								},
								"category": map[string]interface{}{
									"type":        "string",
									"description": "Категория блюда (например: Соусы, Основные блюда, Закуски, Комбо сеты и т.д.)",
								},
							},
							"required": []string{"query", "category"},
						},
					},
				},
				"required": []string{"items"},
			},
		},
	},
	{
		Type: openai.ToolTypeFunction,
		Function: &openai.FunctionDefinition{
			Name:        "set_order",
			Description: "Фиксирует заказ клиента по выбранным позициям меню, вызывается только когда заказ 100% окончен",
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"items": map[string]interface{}{
						"type":        "array",
						"description": "Список заказанных клиентом блюд",
						"items": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"name": map[string]interface{}{
									"type":        "string",
									"description": "Название блюда или позиции из меню",
								},
								"description": map[string]interface{}{
									"type":        "string",
									"description": "Описание блюда или позиции из меню (если нет то продублировать название)",
								},
								"price": map[string]interface{}{
									"type":        "integer",
									"description": "Цена одной порции блюда в тенге",
								},
								"amount": map[string]interface{}{
									"type":        "integer",
									"description": "Количество порций, заказанных клиентом",
								},
							},
							"required": []string{"name", "description", "price", "amount"},
						},
					},
				},
				"required": []string{"items"},
			},
		},
	},
}
