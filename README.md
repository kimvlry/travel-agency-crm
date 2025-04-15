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
  - Дата и время  
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
    - Дата и время  
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
          - Время отправления 
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
  - Дата и время 
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
    - Дата и время 
    - Предпочтительный способ связи  
    - Тема

Все метки времени фиксируются в UTC.


# ERD
![](https://www.plantuml.com/plantuml/png/pLbVRzms4d_Nfo3waZYN0Bq2ULaKHVUIpLrmZfDpxUBUYB4qgoLCaIf9kTd6_kudKMEkQ4ggBhi5yc8lHwFvyvlXa3FwsJemhgbaPc-5gi50nSw5GhUlaMdZTZgJk7NCQMP4kNEi40Pp9xJAiXhoMoYHVPTBWSfznxvaZA5omkrPJx8TXR6_CL6isFkBgyq_DvVinny__VkddvX_ykgL_xTjfEGAAbonNzUNR_wrldp-m-lNP-pYunMxk3u_zqprJgiXm_TFdcC5Giu9k16jFxm0rqfvkxxQn4_prcmkYeE0lt__L8Em_4QsJajX7RQS__ZuyNoplc1lD-_Mr-TNxDtw_9TDo_cG7O2u0MNP5pQ9mm060et7JEywi68_G763OBLQiQlDVwzOhBBpH62hT97EFqCzruroPizZFaZudAZmZoX0hoFu8WjhiBRMXlBZ25oUKvwTURukQLT1EE-MDhPIDTKDcaaR6SFFjJ3GvVTaeWXh6vnyIYPVNxp_p_NcUR3yfJVkBCOXEAOS6l2xohBJGjAPoHkNjvdo_iFcbwlrXqzFurjLZH9vxtE-0wLGhiR9Jy2i49rG9PUwPulJQ12HWNzZGR5DLO7PzydiAL0QnAhpOa0qM0bLe9bncYlyx7YkbMkrySSNR0Bx3FEf0r0RtA8nMF2ZaBOMIgJj7eE7gg1aJ2CyxdUDnce5aYljn9OqVujkEwng2Gv7rGsDvu-Szkp_1cLSpZ8uAByZ4HrfF7n1-Y2fQwCh7VaNczSpZ5Q6CNE4atzIpPL3vP942fPuXqZfX3UvGN1OS71z_UiG3aVWc-lBoyt55G-mUtSU1fU56wrlXIePgCAROo2Fhmx4O1UugvjG3K9UTCJ7D73W6hjwbCZxDTfWxRF2mDOzwvY3BNpaZA1otkC4-wuoSIdKBPMdhuE5iQOkqbS-Nh3hJzrg0ii4A1wgvP8KYJVm-EObmGaJqOdBYbBXucq6rlFZ2MoC9OV0JM_U-MjFUjMZZHsdmD9zT0TI5ArTNPaLkiDdRxqLWMt2_WLc1nGJREW1XOQGIDbIhCwIEqqQoX3kfnpHBcjbMrJkiCjYGBK3oMiZykxQj7dp_iFwlCNzvG-JbyIGrfmadm8WSZY0FebGYkJyfNmUkrc8KjU9x_GP7IJ7TPtD9NLaMbJ-bfYrh9OqPhPp-ytmN1Sumq8GptGV3zdW4EeEYUZGwGZsk0CsEYIwFEVOc2JN8xSQxS2hYkaLWhG-H6mrMDZ79bI2TPLGZSE97hIFvu6DGrqRVJVeMQcHjCYBXlgTwKBrZ5tR1gGac4oZcCKxD22PSDY_Rbahg3lNtGxP7i74sF81unTsYfFTkCaU35Mni3yZFIyF0fCUpRCOtT0kFD9mRGeqlDPMT5ZCdWjUPyig5bK-OazL1-lF5lNxK1G6hPql0iw0wdjiY-PEvCjETLtLeFPpoYjTe8mPnZHlqPmY5i6BX2drZxffhC6unb1OLip_QryFRWGHW-ZHos2Ck0FvDQyUz3fHuU2o6A4NTdB7nFhZwlBNzVbK5BKfGOa_be7vbqxdOhlmSzuOAxGwGP23ambJPDsyT-ceWwNncut1KnU5iYsZof7tkqBWDnQ3OE09Of1ukv-DHQwxewL4waN3otvjPmxrdNOeerFTKxZHkk8vE2orMGQqVo_no5FjsPBR5bq0QZHSTMFpsMj01PzvsQ1rytq2LAY4EpPJecjAWJOtedRYSPHCAYFyZjCtVwe4f0oeSXRGl4i1AMbpeQbIOQdfoMJrwBY7UMfSWuyVfvvsLZewscIhuGmrTc_Xv7IEzsln_KR6fAsRpzZcule3UujReOJFcInJsY5z-jBRiO5z-_AB5-FMlNYnobhV-xhiVpcKM1ge_EBt7JWBTUqNlZRxNtJMjB-zwAS94QGMmkQDjItTBqCU-5LNQ2M8TeMgO9qJJroX6L_mW9O79LwCmj-vB-H-rOPli48mSi-fZvjKGyzx9V4Wxk2A7yFrKBLoF8SeLSjZCA9IyzdJm3ZHDIKjLHIUj6sYPLGo3FxMbLWSFi9Yqknq-E0Te4KSY31FBy21fmp6BARgqO9EyL160a9fGFNs1OVzRuDt-dR6rBJI7PCt4aGrgSe7tqiFNMb8bHk3SBi5sylPWToskG7c41bQJiYUcCDxPU5HualurkU_s-rxNR8J6UgIfr7Yn4J_FRY0FVUvFsLOk82GCM4PULzXqHee_M-fhUEoAKieSKfgi9ns7wqYcJw2YeHkjK5Hge6yhbvTOf_UTYTgcsKq8Rs_V_LAVnce6D3lx-dxSCmyMjc8WpteZYLivK2dLsyWyCO7BkcR_ToMpNwKUj1UmaHkn1xC2Sh69P25mrgOaPS45gMy_qjw1YrrJ8_77Dc0Zv2AMlhuNMfhOsfeX69oo40Yz-6XoreMNWhUHLgJnuCsC0jcqBEHHZyBsWFJI3EJPTwQuFX8ZKf996_saf7AJI_usIYLiZXSewDG-oVfzvGiJ8r9OQZiWqTncjiicR6dBmUUy6Zal4c6m3EyiIVna3Rx6LNHLF9_)