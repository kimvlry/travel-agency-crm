_CRM для турагенства: туроператор занимается созданием авторского (экскурсионная программа включена) или пакетного (перелет-проживание) тура и внесением его в систему, а турагент взаимодействует с клиентами - реализует туры. _

# Функциональные требования
### Клиент:
- Регистрация клиентов с указанием ФИО,номера телефона, даты рождения, города вылета. В дальнейшем (при заключении договора) вносятся паспортные данные.
- История взаимодействий: звонки, письма, бронирования, жалобы, фиксируются с датой и временем.
- Возможность внесения клиента в черный список
- Сегментация клиентов по тегам (например, "VIP", "семейные туры", "экстремальный отдых"). 
- Автоматические уведомления через бот (с кастомизацией - например вставка имени клиента в шаблон): за 6 месяцев до истечения срока паспорта, за сутки до вылета, персональные предложения ко дню рождения, напоминания об оплате за 3 дня до истечения срока оплаты, поздравление с днем рождения
- Рекламная рассылка через бот: анонсы новых туров, горящие туры, начало раннего бронирования

### Тур
- Создание туров с детальным описанием: маршрут, даты, стоимость, включенные услуги.
- Привязка к отелям (номера, условия бронирования, контакты менеджеров отеля), может быть использована для персонализации рассылки.

### Отель
- Запись отелей-партнеров с информацией о номерах, ценах, условиях отмены бронирования.
- История взаимодействий: переписка, согласованные условия фиксируются с датой и временем.

### Финансы и отчетность
- Онлайн-оплата через интеграцию с платежными системами.
- Управление статусами бронирования: «заключение договора», «договор подписан», «частичная оплата», «полная оплата», «запрос отмены бронирования», «бронь аннулирована» 
- Формирование финансовых отчетов: выручка по туроператорам за период, сумма выплаченных комиссий турагентам за период, расходы на отели за период, выписка по всем операциям за период, количество проданных туров за период.
- Отчеты по ключевым метрикам: (конверсия из реакции в соцсети в бронирование, средний чек по направлениям)
- Разграничение доступа: турагент - доступ только к своим турам и клиентам, администратор - полный доступ к данным, бухгалтер - доступ к фин.отчетам
- Шифрование персональных данных клиентов (паспорта, платежные реквизиты).
- Возможность создания договора и отправка туристу доступа на прочтение и подписание в электронном виде.
- Автоматическое формирование согласия на обработку персональных данных клиенту на подпись, формирование ссылки или qr-кода на оплату

