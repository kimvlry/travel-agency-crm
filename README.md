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
left to right direction

package "client" {
  entity clients{
    + id: INTEGER <<PK>> 
    --
    full_name: VARCHAR(255) NOT NULL
    phone: VARCHAR(20) NOT NULL
    email: VARCHAR(255) NOT NULL
    birth_date: DATE NOT NULL
    city_id: VARCHAR(100) NOT NULL
    is_blacklisted: BOOLEAN DEFAULT FALSE
  }

  entity bans { 
    + id: INTEGER <<PK>>
    --
    client_id: INTEGER <<FK>> NOT NULL
    ban_reason: TEXT 
  }

  entity cities {
    + id: INTEGER <<PK>>
    --
    name: VARCHAR(255) NOT NULL
    country_name: VARCHAR(a55) NOT NULL
    timezone: VARCHAR(50) NOT NULL
  }

  entity passports{
    + id: INTEGER <<PK>>
    --
    client_id: INTEGER <<FK>> (clients)
    type: passport_types NOT NULL
    number: VARCHAR(50) NOT NULL
    expiration_date: DATE NOT NULL
    issue_date: DATE NOT NULL
    --
    UNIQUE(client_id, type)
  }

  entity client_interactions {
    + id: INTEGER <<PK>>
    --
    client_id: INTEGER <<FK>> (clients)
    date_utc: TIMESTAMP NOT NULL
    communication_channel: communication_channels NOT NULL
    meeting_location: VARCHAR(255)
    type: interaction_types NOT NULL
    summary: TEXT
    agreements: TEXT
    reminder_id: INTEGER <<FK>> (client_next_contact_reminders)
  }

  entity client_next_contact_reminders {
    + id: INTEGER <<PK>>
    --
    client_id: INTEGER <<FK>> (clients)
    preferred_communication_channel: communication_channels NOT NULL
    message: TEXT NOT NULL
    send_date: TIMESTAMP NOT NULL
  }

  entity client_personal_notification {
    + id: INTEGER <<PK>>
    --
    client_id: INTEGER <<FK>> (clients)
    preferred_communication_channel: communication_channels NOT NULL
    template_id: INTEGER <<FK>> (notification_templates)
    send_date: TIMESTAMP NOT NULL
  }

  entity notification_templates {
    + id: INTEGER <<PK>>
    --
    type: notification_types NOT NULL
    message_template: TEXT
    promo_id: INTEGER 
  }

  entity promotions {
    + id: INT <<PK>>
    --
    title: VARCHAR(255) NOT NULL
    content: TEXT NOT NULL
    promo_type: promotion_types NOT NULL
    created_at: DATETIME DEFAULT CURRENT_TIMESTAMP
  }
}

package "booking and contract" {
  entity bookings{
    + id: INT <<PK>>
    --
    tour_id: INT <<FK>> (tours)
    status: booking_statuses DEFAULT 'draft'
    contract_number: VARCHAR(50) UNIQUE
    payment_links: TEXT
    created_at: DATETIME DEFAULT CURRENT_TIMESTAMP
    updated_at: DATETIME ON UPDATE CURRENT_TIMESTAMP
  }

  entity booking_agreements {
    + id: INT <<PK>>
    --
    client_id: INT <<FK>> (clients)
    booking_id: INT <<FK>> (bookings)
    signed_date: TIMESTAMP NOT NULL
  }

  entity assignees {
    + id: INT <<PK>>
    --
    client_id: INT <<FK>> (clients)
  }

  entity contract_templates {
    + id: INT <<PK>>
    --
    name: VARCHAR(100) NOT NULL
    type: VARCHAR(50) NOT NULL
    content: TEXT NOT NULL
    validity_period_days: INT NOT NULL
  }

  entity contracts {
    + id: INT <<PK>>
    --
    assignee_id: INT <<FK>> (assignes)
    template_id: INT <<FK>> (contract_templates)
    issue_date: TIMESTAMP NOT NULL
    sign_date: TIMESTAMP
    status: consent_statuses NOT NULL
    total_price: DECIMAL(10,2) NOT NULL
  }

  entity agreement_consents {
    + id: INT <<PK>>
    --
    assignee_id: INT <<FK>> (assignees)
    contract_id: INT <<FK>> (contracts)
    template_id: INTEGER <<FK>> NOT NULL
    date: TIMESTAMP NOT NULL
    status: consent_statuses NOT NULL
  }

  entity consent_templates {
    + id: INT <<PK>>
    --
    name: VARCHAR(100) NOT NULL
    type: consent_types NOT NULL 
    content: TEXT NOT NULL
  }

  entity payment_links {
    + id: INT <<PK>>
    --
    booking_id: INT <<FK>> (bookings)
    url: VARCHAR(255) NOT NULL
    qr_code: VARCHAR(255) NOT NULL
    status: VARCHAR(50) NOT NULL
  }
}


package "tour" {
  entity tours{
    + id: INT <<PK>>
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
    + id: INT <<PK>>
    --
    tour_id: INT <<FK>> (tours)
    start_date: DATE NOT NULL
    end_date: DATE NOT NULL
    --
    UNIQUE(tour_id, start_date)
  }

  entity tour_routes {
    + id: INT <<PK>>
    --
    tour_id: INT <<FK>> (tours)
    order_position: INT NOT NULL
  }

  entity route_points {
    + id: INT <<PK>>
    --
    route_id: INT <<FK>> (tour_routes)
    name: VARCHAR(255) NOT NULL
    address: TEXT NOT NULL
  }

  entity transport_services {
    + id: INT <<PK>>
    --
    company: VARCHAR(255) NOT NULL
    model: VARCHAR(255)
  }

  entity transfers {
    + id: INT <<PK>>
    --
    tour_id: INT <<FK>> (tours)
    transport_id: INT <<FK>> (transport_services)
    departure_point: point_id NOT NULL <<FK>> (route_points)
    arrival_point: point_id NOT NULL <<FK>> (route_points)
    departure_time: DATETIME NOT NULL
    duration_time : INTERVAL NOT NULL
  }

  entity organizers {
    + id: INT <<PK>>
    --
    name: VARCHAR(255) NOT NULL
    phone: VARCHAR(20) NOT NULL
    email: VARCHAR(255) NOT NULL
  }

  entity excursions {
    + id: INT <<PK>>
    --
    tour_id: INT <<FK>> (tours)
    organizer_id: INT <<FK>> NOT NULL
    name: VARCHAR(255) NOT NULL
    meeting_location : point_id NOT NULL
    duration_time : INTERVAL
  }

  entity insurance_companies {
    + id: INT <<PK>>
    --
    name: VARCHAR(255) NOT NULL
    phone: VARCHAR(20) NOT NULL
    email: VARCHAR(255) NOT NULL
  }

  entity insurances {
    + id: INT <<PK>>
    --
    tour_id: INT <<FK>> (tours)
    insurance_company_id: INT <<FK>> NOT NULL
    coverage_type: insurance_types NOT NULL
  }
}


package "hotel" {
  entity hotel_room_categories {
    + id: INT <<PK>>
    --
    hotel_id: INT <<FK>> (hotels)
    name: VARCHAR(100) NOT NULL
    price_per_night: DECIMAL(10,2) NOT NULL
    max_guests: INT NOT NULL
  }

  entity amenities {
    + id: INT <<PK>>
    --
    name: VARCHAR(255) NOT NULL UNIQUE
    description: TEXT
  }

  entity hotels {
    + id: INT <<PK>>
    --
    name: VARCHAR(255) NOT NULL
    address: TEXT NOT NULL
    price_range: VARCHAR(100) NOT NULL
    cancellation_terms: TEXT NOT NULL
  }

  entity hotel_interactions {
    + id: INT <<PK>>
    --
    hotel_id: INT <<FK>> (hotels)
    date_utc: DATETIME NOT NULL
    communication_channel: communication_channels NOT NULL
    type: interaction_types NOT NULL
    summary: TEXT
    agreements: TEXT
    next_contact_reminder : reminder_id <<FK>> (hotel_next_contact_reminders)
  }

  entity hotel_next_contact_reminders {
    + id: INT <<PK>>
    --
    hotel_id: INT <<FK>> 
    preferred_communication_channel: communication_channels NOT NULL
    message: TEXT NOT NULL
    send_date: TIMESTAMP NOT NULL
  }
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

note top of passports
  **passport_types**
  - internal
  - foreign
end note



' Relationships

clients||--o{ passports
clients||--|| cities
clients||--o{ client_interactions
clients }o--o{ bookings
clients ||--o{ bans
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
excursions ||--|| organizers
transfers ||--|| transport_services
insurances ||--|| insurance_companies

hotels ||--o{ hotel_interactions
hotel_interactions ||--|| hotel_next_contact_reminders
hotels ||--o{ hotel_room_categories
hotel_room_categories }o--o{ amenities

bookings||--o{ booking_agreements
assignees ||--o{ contracts
assignees ||--|| clients
contract_templates ||--o{ contracts
contracts ||--o{ agreement_consents
consent_templates ||--o{ agreement_consents
bookings||--o{ payment_links

@enduml
```