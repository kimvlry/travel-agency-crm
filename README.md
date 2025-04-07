# CRM для турагенства, специализирующегося на создании и реализации авторских туров 
Две ключевые роли в таком турагенстве - туроператор (разрабатывает туры, продумывает маршруты, экскурсии, логистику. Согласовывает условия с партнерами: отелями, транспортными компаниями)
и менеджер (работает с клиентами: консультирует, поддерживает связь на всех этапах, контролирует оплату и документооборот)

## Работа менеджера
#### Взаимодействие с клиентом
- Система должна регистрировать клиента с обязательными полями:  
  - ФИО  
  - Номер телефона  
  - Email
  - Дата рождения  
  - Город вылета  
- Система должна предоставлять возможность внесения паспортных данных клиента при заключении договора.  

- Система должна фиксировать все контакты с клиентом, включая:  
  - Дата и время в формате UTC  
  - Средство связи:  
    - Телефонный звонок 
    - Telegram  
    - WhatsApp  
    - Электронная почта  
    - Личная встреча (с указанием места)  
  - Тип взаимодействия:  
    - Обсуждение (выявление потребностей, уточнение деталей)  
    - Согласование условий  
    - Рекламация (претензия к услугам)  
  - Итоги:  
    - Краткое описание (1-2 предложения)  
    - Зафиксированные договорённости  
  - Напоминание о следующем шаге:  
    - Дата и время (с учётом часового пояса клиента)  
    - Предпочтительный способ связи  
    - Тема

- Система должна предоставлять возможность добавления клиента в черный список
- Система должна предоставлять возможность для автоматической рассылки уведомлений клиентам:  
    - за 6 месяцев до истечения срока паспорта
    - за сутки до вылета
    - напоминания об оплате за 3 суток до истечения срока оплаты
  
- Система должна предоставлять возможность для автоматической рекламной рассылки: 
    - анонсы новых авторских туров, предлагаемых агенством
    - горящие туры ([подробнее](#создание-тура))
    - начало раннего бронирования 
    - поздравление и персональное предложение в 12:00 по локальному времени клиента в день его рождения

#### Управление бронированием
- Система должна позволять привязывать к одному бронированию список клиентов, зарегистрированных в системе

- Система должна поддерживать следующие статусы договора бронирования: 
    - черновик
    - ожидает подписания договора
    - ожидает оплаты
    - частичная оплата
    - полная оплата
    - запрос отмены бронирования
    - аннулирование 

- Система должна автоматически:  
  - Генерировать договор, подставляя в шаблон:  
    - Данные клиента (ФИО, паспорт)  
    - Параметры тура (даты, стоимость, услуги)  
  - Отправлять договор клиенту через Контур.Крипто для подписания.  

- Система должна автоматически:  
  - Создавать согласие на обработку персональных данных по шаблону.  
  - Отправлять согласие клиенту вместе с договором через Контур.Крипто.
  
- Система должна:  
  - Интегрироваться с платежной системой ЮKassa.  
  - Формировать ссылку/QR-код для оплаты бронирования.  
  - Обновлять статус договора при успешной оплате.  

## Работа туроператора
#### Создание тура 
- Система должна регистрировать тур с обязательными параметрами:
    - Уникальное название тура
    - Даты: согласованные с отелями и перевозчиками интервалы дат на 6–12 месяцев вперед
    - **Маршрут** - последовательность точек тура в порядке посещения
      - Последовательностью точек (с возможностью изменения порядка)  
      - Для каждой точки:
        - Название
        - Адрес  
        - Длительность пребывания в формате (дни-часы-минуты)
    - Стоимость - конечная стоимость для туриста (EUR)
    - Включенные услуги 
        - **Трансфер**:  
          - Точки отправления/прибытия (с привязкой к карте)  
          - Время в пути (дни-часы-минуты)  
          - Время отправления (с учётом часового пояса клиента)  
          - Транспортная компания (выбор из базы партнёров)  
          - Модель транспортного средства  

        - **Экскурсии**
            - Название
            - Контакты организатора - имя, телефон, email

        - **Тип питания** 
          - Завтраки
          - Полупансион
          - Полный пансион
          
        - **Включенная страховка**
            - Тип покрытия
              - Медицинская
              - Потеря багажа 
              
            - Страховая компания - название, телефон, email

- Система должна привязывать зарегистрированные в системе отели к турам

### <span id="hidden-anchor"></span>
- Система должна автоматически помечать тур как "горящий" при:  
  - Продаже >60% мест от общей квоты (не требует подтверждения статуса)
  - Остатке ≤30 дней до начала тура (требует ручного подтверждения статуса менеджером)

#### Взаимодействие с отелем
- Система должна хранить записи отелей-партнеров, включая:  
   - Адрес
   - Перечень категорий номеров (например: стандартный, люкс, семейный)
   - Диапазон цен за категорию номера (EUR).  
   - Условия отмены бронирования
     - Сроки
     - Штрафы
     - Исключения

- Система должна фиксировать все контакты с администрацией отеля, включая:  
  - Дата и время в формате UTC  
  - Средство связи:  
    - Телефонный звонок 
    - Telegram  
    - WhatsApp  
    - Электронная почта  
    - Личная встреча (с указанием места)  
  - Тип взаимодействия:  
    - Обсуждение 
    - Согласование условий  
    - Рекламация 
  - Итоги:  
    - Краткое описание (1-2 предложения)  
    - Зафиксированные договорённости  
  - Напоминание о следующем шаге:  
    - Дата и время (с учётом часового пояса, в котором расположен отель)  
    - Предпочтительный способ связи  
    - Тема


# ERD
```plantuml
@startuml
skinparam Linetype ortho


' ----- clients-----

entity clients{
  + client_id: INTEGER <<PK>> 
  --
  full_name: VARCHAR(255) NOT NULL
  phone: VARCHAR(20) NOT NULL
  email: VARCHAR(255) NOT NULL
  birth_date: DATE NOT NULL
  city_id: VARCHAR(100) NOT NULL
  is_blacklisted: BOOLEAN DEFAULT FALSE
}

entity bans { 
  + ban_id: INTEGER <<PK>>
  --
  client_id: INTEGER <<FK>> NOT NULL
  ban_reason: TEXT 
}

entity cities {
  + city_id: INTEGER <<PK>>
  --
  city_name: VARCHAR(255) NOT NULL
  country_name: VARCHAR(a55) NOT NULL
  timezone: VARCHAR(50) NOT NULL
}

entity passports{
  + passport_id: INTEGER <<PK>>
  --
  client_id: INTEGER <<FK>> (clients)
  number: VARCHAR(50) NOT NULL
  expiration_date: DATE NOT NULL
  issue_date: DATE NOT NULL
}

entity foreign_passports{
  + passport_id: INTEGER <<PK>>
  --
  client_id: INTEGER <<FK>> (clients)
  number: VARCHAR(50) NOT NULL
  expiration_date: DATE NOT NULL
  issue_date: DATE NOT NULL
}

entity client_interactions {
  + interaction_id: INTEGER <<PK>>
  --
  client_id: INTEGER <<FK>> (clients)
  interaction_date_utc: TIMESTAMP NOT NULL
  communication_channel: communication_channels NOT NULL
  meeting_location: VARCHAR(255)
  interaction_type: interaction_types NOT NULL
  summary: TEXT
  agreements: TEXT
  reminder_id: INTEGER <<FK>> (client_next_contact_reminders)
}

entity client_next_contact_reminders {
  + reminder_id: INTEGER <<PK>>
  --
  client_id: INTEGER <<FK>> (clients)
  preferred_communication_channel: communication_channels NOT NULL
  message: TEXT NOT NULL
  send_date: TIMESTAMP NOT NULL
}

entity client_personal_notification {
  + notification_id: INTEGER <<PK>>
  --
  client_id: INTEGER <<FK>> (clients)
  preferred_communication_channel: communication_channels NOT NULL
  template_id: INTEGER <<FK>> (notification_templates)
  send_date: TIMESTAMP NOT NULL
}

entity notification_templates {
  + template_id: INTEGER <<PK>>
  --
  type: notification_types NOT NULL
  message_template: TEXT
  promo_id: INTEGER 
}



' ----- BOOKING & CONTRACTS -----

entity bookings{
  + booking_id: INT <<PK>>
  --
  tour_id: INT <<FK>> (tours)
  status: booking_statuses DEFAULT 'draft'
  contract_number: VARCHAR(50) UNIQUE
  payment_links: TEXT
  created_at: DATETIME DEFAULT CURRENT_TIMESTAMP
  updated_at: DATETIME ON UPDATE CURRENT_TIMESTAMP
}

entity client_bookings {
  + client_id: INT <<PK,FK>> (clients)
  + booking_id: INT <<PK,FK>> (bookings)
}

entity booking_agreements {
  + agreement_id: INT <<PK>>
  --
  client_id: INT <<FK>> (clients)
  booking_id: INT <<FK>> (bookings)
  signed_date: TIMESTAMP NOT NULL
}

entity assignees {
  + assignee_id: INT <<PK>>
  --
  client_id: INT <<FK>> (clients)
}

entity contract_templates {
  + template_id: INT <<PK>>
  --
  name: VARCHAR(100) NOT NULL
  type: VARCHAR(50) NOT NULL
  template_content: TEXT NOT NULL
  validity_period_days: INT NOT NULL
}

entity contracts {
  + contract_id: INT <<PK>>
  --
  assignee_id: INT <<FK>> (assignes)
  contract_template_id: INT <<FK>> (contract_templates)
  issue_date: TIMESTAMP NOT NULL
  sign_date: TIMESTAMP
  status: consent_statuses NOT NULL
  total_price: DECIMAL(10,2) NOT NULL
}

entity agreement_consents {
  + consent_id: INT <<PK>>
  --
  assignee_id: INT <<FK>> (assignees)
  contract_id: INT <<FK>> (contracts)
  consent_template_id: INTEGER <<FK>> NOT NULL
  consent_date: TIMESTAMP NOT NULL
  status: consent_statuses NOT NULL
}

entity consent_templates {
  + consent_template_id: INT <<PK>>
  --
  name: VARCHAR(100) NOT NULL
  type: consent_types NOT NULL 
  template_content: TEXT NOT NULL
}

entity payment_links {
  + payment_link_id: INT <<PK>>
  --
  booking_id: INT <<FK>> (bookings)
  payment_url: VARCHAR(255) NOT NULL
  qr_code: VARCHAR(255) NOT NULL
  status: VARCHAR(50) NOT NULL
}



' ----- TOUR -----

entity tours{
  + tour_id: INT <<PK>>
  --
  title: VARCHAR(255) NOT NULL UNIQUE
  price_eur: DECIMAL(10,2) NOT NULL
  quota: INT NOT NULL
  meals_type : meals_types
  is_lastminute: BOOLEAN DEFAULT FALSE 
  lastminute_approved: BOOLEAN
  base_duration_days: INT NOT NULL
}

' Usually toursrun several iterations during the year
entity tour_iterations { 
  + schedule_id: INT <<PK>>
  --
  tour_id: INT <<FK>> (tours)
  start_date: DATE NOT NULL
  end_date: DATE NOT NULL
  --
  UNIQUE(tour_id, start_date)
}

entity tour_routes {
  + route_id: INT <<PK>>
  --
  tour_id: INT <<FK>> (tours)
  order_position: INT NOT NULL
}

entity route_points {
  + point_id: INT <<PK>>
  --
  route_id: INT <<FK>> (tour_routes)
  name: VARCHAR(255) NOT NULL
  address: TEXT NOT NULL
}

entity transfers {
  + transfer_id: INT <<PK>>
  --
  tour_id: INT <<FK>> (tours)
  departure_point: point_id NOT NULL <<FK>> (route_points)
  arrival_point: point_id NOT NULL <<FK>> (route_points)
  transport_company: VARCHAR(255) NOT NULL
  vehicle_model: VARCHAR(255)
  departure_time: DATETIME NOT NULL
  duration_time : INTERVAL NOT NULL
}

entity excursions {
  + excursion_id: INT <<PK>>
  --
  tour_id: INT <<FK>> (tours)
  name: VARCHAR(255) NOT NULL
  meeting_location : point_id NOT NULL
  organizer_name: VARCHAR(255) NOT NULL
  organizer_phone: VARCHAR(20) NOT NULL
  organizer_email: VARCHAR(255) NOT NULL
  duration_time : INTERVAL
}

entity insurances {
  + insurance_id: INT <<PK>>
  --
  tour_id: INT <<FK>> (tours)
  coverage_type: insurance_types NOT NULL
  company_name: VARCHAR(255) NOT NULL
  company_phone: VARCHAR(20) NOT NULL
  company_email: VARCHAR(255) NOT NULL
}

entity hotels {
  + hotel_id: INT <<PK>>
  --
  name: VARCHAR(255) NOT NULL
  address: TEXT NOT NULL
  room_categories: JSON NOT NULL
  price_range: VARCHAR(100) NOT NULL
  cancellation_terms: TEXT NOT NULL
}

entity hotel_interactions {
  + interaction_id: INT <<PK>>
  --
  hotel_id: INT <<FK>> (hotels)
  interaction_date_utc: DATETIME NOT NULL
  communication_channel: communication_channels NOT NULL
  interaction_types: interaction_types NOT NULL
  summary: TEXT
  agreements: TEXT
  next_contact_reminder : reminder_id <<FK>> (hotel_next_contact_reminders)
}

entity hotel_next_contact_reminders {
  + reminder_id: INT <<PK>>
  --
  hotel_id: INT <<FK>> 
  preferred_communication_channel: communication_channels NOT NULL
  message: TEXT NOT NULL
  send_date: TIMESTAMP NOT NULL
}

entity promotions {
  + promo_id: INT <<PK>>
  --
  title: VARCHAR(255) NOT NULL
  content: TEXT NOT NULL
  promo_type: promotion_types NOT NULL
  created_at: DATETIME DEFAULT CURRENT_TIMESTAMP
}



' ENUM Definitions

note top of client_interactions
  **communication_channels**:
  - phone
  - telegram
  - whatsapp
  - email
  - meeting
  
  **interaction_types**:
  - discussion
  - agreement
  - complaint
end note

note top of promotions
  **promotion_types**
  - new_tours
  - hot_tours
  - early_booking
end note

note top of bookings
  **booking_statuses**:
  - draft
  - pending_signature
  - pending_payment
  - partially_paid
  - fully_paid
  - cancellation_requested
  - canceled
end note

note top of agreement_consents
  **consent_types**:
  - personal_data
  - contract_terms
  - ad
  
  **consent_statuses**:
  - granted
  - pending
  - revoked
end note

note top of hotel_interactions
  **interaction_types**:
  - discussion
  - agreement
  - claim
end note

note top of tours
  **meals_types**
  - breakfast
  - half_board
  - full_board
end note

note top of notification_templates
  **notification_types**:
  - passport_expiry
  - flight_reminder
  - payment_reminder
  - birthday_promo
end note

note top of insurances
  **insurance_types**
  - medical
  - lost_luggage
end note



' Relationships

clients||--o{ passports
clients||--o{ foreign_passports
clients||--|| cities
clients||--o{ client_interactions
clients||--o{ client_bookings
clients ||--o{ bans
client_bookings }o--|| bookings
client_interactions ||--|| client_next_contact_reminders
client_personal_notification }o--|| notification_templates
clients||--o{ client_personal_notification
notification_templates ||--|| promotions

bookings||--|| tours
tours||--o{ tour_routes
tours||--o{ tour_iterations
tour_routes ||--o{ route_points
tours||--o{ transfers
tours||--o{ excursions
tours||--o{ insurances
tours}o--o{ hotels
excursions ||--|| route_points

hotels ||--o{ hotel_interactions
hotel_interactions ||--|| hotel_next_contact_reminders

bookings||--o{ booking_agreements
assignees ||--o{ contracts
assignees ||--|| clients
contract_templates ||--o{ contracts
contracts ||--o{ agreement_consents
consent_templates ||--o{ agreement_consents
bookings||--o{ payment_links

@enduml
```