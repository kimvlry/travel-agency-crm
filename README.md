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
![](https://www.plantuml.com/plantuml/svg/p5bVRzkm4N_tfo3waZXhWAv0NuoYcDUwMp37wLAx6FP2NAIpJ8GYLP9Au_xvxfCeqX5jqL4RTUYJJV9uVtvtF1vFVzC6b2bpFjAtJ1IW82Tp9j1i2oHIcOqSSLmROYHHBDiOaZA5YM5IZ4O59BUG8NcMS8R2F2FVHeIWCCniIJEbwnb2_a9OEY4NY-Ni7xDhylhrXt-zUKFiolcv_LcNd5C1EKx8f-drstzEhqzVlder9ekh9LcivdDBL6oaQ1Eys5l771W_nk26LVRG54p5vTrqEGjNaqfjoj871dzziI-1QNh3AwCvqmOhohzVNSrdqmLvDti_NSsNvFrq_d5MKVuOFG1n0qAJRoIAGmi61neDYTvNH7jsWA0AGKin8SlPVvPaJsGYIs4KmqfkBx77WDzdpSmWVErwe-SsiDSJUCDFTyQCxHx3Slmg1LAvNciqyuEDUmeNe7KXbOt49tdWj5doQcmBdEoOqteOsYRA_0RLW_Ml3gBrlc0AweCKZKYcTOdHLQVnQd7nxzNiTAVuSwlRU0y6lomCAh37LvE14M6rivSNbxEFo-dbX-L_ztsTvwLWIMDniW4Xa4-wfyETEQ9X8gDSDbHXHBMyqJAkmo6wp7DGs-R4s1d8569UMz2QL9WpaQ8wOZ0LU6ze8eMfX56_GHz27YSU6ln2uHgLmfG-1MUj8SC6YptiKAGqxk1kmmjKMWhWL4Z3raxq7sYtmRpWOB1JMbjtwYdr-1TGwURK4myRu26BUghJWJlchQWkbCnbO62eNKFmCsd1CCCVIUx2e327uUIrQMoo_oC69Gh1O4h1X0Zxw_RjwlfwjbZItPerwaUhCBcHyfQ9Z818hK8Aah1CSGIwfz6ox4uCzOAF2WEcr1FFcJRZdSudgOArER6aNYFQSLculDxW1TiwJL7EnAtBLR-C3o5baNRkkbgGrGTxpNJZUW0QtINGeTEP4z19wDOysYmJcDexijSf16rtu90QXnAyKyCpthzQYXMVDWU48TBtlDq1PscbMvsIcQpHskh3YYjYHr_rFRATMBb5FNuyunxYDmwBeQYZNJWSqWJdikAlQvSsups8fG5E2yMIkkAQlRsudCuhJpn_EOwXj3i2rF4T4YvlkeVaA5uzi0riZPML4PmYzZkgtnJiddizwXNjeN91tkofMF_KKwgZZytFYYOonJY9XpfsfkrLrhhBwZickBlg2VtqsxfzqzZGfrYgQFXRqqffO2TklmW1hXk7JLe3_tpce0tDcIWDnf_FXBJ8A1I5adUjvxPx0skaQUbUKCSpsWbPwH8uTv0HLGgYyGuLSC8CDZmqgRZLjOBP8DaYg11esYByHeQi5fI9Fl5Gf3sVVqxMynR3_Ol9aYXPsaCwY04XVykQ5fBrJu9sIrI2KtOST6SEADjYnxtQ3f2c2hNkkYbzGDcodZJ2hZzDvn5j_TwusamqGofLYgfIKxEQ_wEthr4WcjQ2HdN7alu5YymB4DjZzkSoHHuIT4bVenekK4ZBeYx20tlThXGBKAPKpiCJOdyg5W7saK00fTWTy9_T6iZj_y2JAWF1llP7xRVsEqFTy3ufbMPI3EdHdS77QusufT4M3-djh61xRrSneKi58a7Q71U6VvRFGYK7zDc1vTj7VPV8-gBCqBtNFODcVAHSsKY3FAXNx0nLKkOq0OEPLFrnjtixZRChUjodgdJbJO6AYlfRnj4Y9uTxcfMepUDF9SXH3DCN3ugn57MYM66Ox_CxiG6U0yYCNOeXQWf4TXJSf8uApds3I-M7p0xL3phIWyP22WPfQPAcTggJmW0DwQ6xpPtDOJ9fTw53uthsde_HFX7b_rE_-T6UJcNk2PajLfVa7QwPO3Q2HYCX3RelbNBTzTcZsdnsrgdUsTcaMZnlShZzPv1ZfY2tWoyRC1gAmWvi7hV_tELJ_RUSzyF2CqsPJagjgsaxt4M37TMvc0EhHYXIKfkmRuXhathz_T3Bi5m4Vg4srTjHwmbbPnkrGV4jTU_Og32tteWA-vZU7Dl5RD2gkDXrbecAGc4mwvxWpHmem-fdM3NBKZkrBeDXa48KVgxJBxQNC8qgVTY9iUgxGU2BtPU950pOcLRJIUKDVb3B3LY426GAX5VEMMl_AxoJjt4zEzBU4-A60ykZacmiMERs7TuEbHk5SBi6tR3P05zNGG7gmIbk664T-PXXPGLB0Uh-OwdzyhbjHF6QuoutsJaNCk6a_NIVmfRQo8_gvQiJfyjksBOznvGbmEr_BhMXlCmoo33Ar2lUyFIZDalhFE5uhgL2beaMkpfJNMCJrdh3YYfDkURnz-_dv_9_VPen4a8X34JxFOL7yIfQwrWuLkBTrSyRbamoQJTB0XZoMUBR4NPy3BhjaLpMD1Ydp-EWS6QDC_HwcTpK2YPUsx9KszVxHnwZJ8oqWiwHQW_CSA-3MlyCc3OsIomUb5EmYuP0FypCo-gzmFt-8_FMpkx-urYaT42jnmBiz0NcvvpN6daGkn6ZPhCLgQpeiMN1BKU7LEaaQSw6XLVbB4FuQwybsMyM2ofHijxNSES0pdB3Icu4mU4J44pXM3oLMXDfi3dBL7oNqgDvmuqhYK_QElzgpEj--m80)