# Dokument wymagań produktu (PRD) - EV Chargers Management System

## 1. Przegląd produktu

### Nazwa produktu
EV Chargers Management System (ECMS)

### Wersja dokumentu
1.0

### Data utworzenia
2025

### Opis produktu
EV Chargers Management System to aplikacja webowa służąca do zarządzania fizycznymi stacjami ładowania pojazdów elektrycznych oraz ich portami. System umożliwia tworzenie lokalizacji, w których mogą znajdować się liczne stacje ładowania, oraz automatyczne generowanie punktów ładowania (EVSE) w oparciu o porty stacji.

### Cel biznesowy
Stworzenie centralnego systemu zarządzania infrastrukturą ładowania pojazdów elektrycznych, który umożliwi efektywne zarządzanie stacjami, lokalizacjami i punktami ładowania w ramach jednej platformy.

### Grupa docelowa
- Właściciele stacji ładowania pojazdów elektrycznych
- Operatorzy infrastruktury ładowania
- Administratorzy systemów zarządzania energią

### Stack technologiczny
Backend:
- REST API w Go
- Framework Gin
- GORM do obsługi bazy danych

Frontend:
- SPA w Angular 20 z TailwindCSS

Baza danych:
- PostgreSQL

CI/CD:
- Github Actions


## 2. Problem użytkownika

### Główny problem
Obecnie brakuje zintegrowanego systemu, który umożliwiałby jednoczesne zarządzanie fizycznymi stacjami ładowania pojazdów EV wraz z ich portami oraz tworzenie logicznych lokalizacji, w których może znajdować się kilka stacji. Użytkownicy muszą zarządzać tymi zasobami w sposób rozproszony, co prowadzi do:

- Braku centralnego widoku wszystkich zasobów ładowania
- Trudności w organizacji stacji według lokalizacji fizycznych
- Braku automatycznego generowania punktów ładowania z portów stacji
- Braku jednolitego systemu autoryzacji i audytu operacji

### Problemy szczegółowe
1. **Zarządzanie stacjami**: Brak możliwości centralnego zarządzania stacjami ładowania wraz z ich specyfikacjami technicznymi
2. **Organizacja lokalizacji**: Trudności w grupowaniu stacji według fizycznych lokalizacji
3. **Generowanie punktów ładowania**: Manualne tworzenie punktów ładowania zamiast automatycznego generowania z portów stacji
4. **Kontrola dostępu**: Brak systemu ról i uprawnień dla różnych typów użytkowników
5. **Audyt operacji**: Brak śledzenia zmian i operacji wykonywanych na zasobach

## 3. Wymagania funkcjonalne

### 3.1 Zarządzanie stacjami ładowania
- Tworzenie, odczyt, aktualizacja i usuwanie stacji ładowania
- Zarządzanie portami stacji (ConnectorId, Power, Voltage, Amperage, ConnectorType, ConnectorStandard)
- Status stacji: "w magazynie" lub przypisana do lokalizacji
- Wyszukiwanie stacji po nazwie, modelu lub numerze seryjnym

### 3.2 Zarządzanie lokalizacjami
- Tworzenie lokalizacji fizycznych z danymi adresowymi
- Przypisywanie stacji do lokalizacji
- Ręczne dodawanie stacji do lokalizacji z prostym wyszukiwaniem
- Walidacja kodu kraju zgodnie z ISO 3166-1 alpha-3

### 3.3 Generowanie punktów ładowania (EVSE)
- Automatyczne generowanie punktów ładowania z portów stacji (mapowanie 1:1)
- Generowanie EvseID zgodnego ze specyfikacją Emi3spec
- Powiązanie punktów ładowania z lokalizacjami

### 3.4 System użytkowników i autoryzacji
- Rejestracja i logowanie użytkowników
- Dwie role: Admin (dostęp globalny) i Owner (dostęp tylko do własnych zasobów)
- Autentykacja JWT
- Middleware autoryzacji oparty na rolach

### 3.5 Audyt i logowanie
- Logowanie wszystkich requestów HTTP
- Szczegółowy audyt operacji CRUD z informacją o poprzednich i nowych wartościach
- Śledzenie zmian w zasobach systemu

## 4. Granice produktu

### 4.1 W zakresie MVP
- CRUD operacje dla stacji ładowania i ich portów
- CRUD operacje dla lokalizacji
- Automatyczne generowanie punktów ładowania z portów stacji
- System kont użytkowników z dwoma rolami (Admin, Owner)
- Mapowanie 1:1 między portami stacji a punktami ładowania
- Model single-tenant (jeden użytkownik = jeden tenant)
- Audyt i logowanie operacji
- REST API w Go + SPA w Angular
- Baza danych PostgreSQL

### 4.2 Poza zakresem MVP
- Integracja z zewnętrznymi usługami po protokołach EV Roaming Foundation
- Import masowy stacji oraz lokalizacji
- Udostępnianie stacji i lokalizacji innym podmiotom
- Zarządzanie różnicami modeli stacji (warstwa abstrakcji)
- Płatności i rozliczenia
- Aplikacja mobilna
- Real-time monitoring stacji
- Multi-tenant architecture
- Zaawansowane raportowanie i analityka

## 5. Historyjki użytkowników

### US-001: Rejestracja nowego użytkownika
**Tytuł**: Rejestracja nowego użytkownika w systemie

**Opis**: Jako nowy użytkownik chcę móc zarejestrować się w systemie, aby uzyskać dostęp do zarządzania stacjami ładowania.

**Kryteria akceptacji**:
- Użytkownik może zarejestrować się podając email i hasło
- System waliduje format emaila
- Hasło musi mieć minimum 8 znaków
- Po rejestracji użytkownik otrzymuje rolę "Owner" domyślnie
- System generuje JWT token po udanej rejestracji
- Operacja rejestracji jest logowana w systemie audytu

### US-002: Logowanie użytkownika
**Tytuł**: Logowanie istniejącego użytkownika

**Opis**: Jako zarejestrowany użytkownik chcę móc zalogować się do systemu używając moich danych dostępowych.

**Kryteria akceptacji**:
- Użytkownik może zalogować się podając email i hasło
- System waliduje dane logowania
- Po udanym logowaniu system zwraca JWT token
- Token zawiera informacje o roli użytkownika
- Operacja logowania jest logowana w systemie audytu
- W przypadku błędnych danych system zwraca odpowiedni komunikat błędu

### US-003: Dodanie pierwszej stacji ładowania
**Tytuł**: Dodanie nowej stacji ładowania przez Owner

**Opis**: Jako Owner chcę dodać nową stację ładowania do systemu, aby móc nią zarządzać.

**Kryteria akceptacji**:
- Owner może dodać stację podając: Vendor (string), Model (string), SerialNumber (string)
- System generuje unikalny UUID dla stacji
- Stacja otrzymuje domyślny status "w magazynie"
- Owner może dodać porty do stacji z następującymi danymi:
  - ConnectorId (int)
  - Power (float32)
  - Voltage (int)
  - Amperage (int)
  - ConnectorType (enum: CCS/Type2/Chademo)
  - ConnectorStandard (enum: AC_1P/AC_3P/DC)
- Operacja dodania stacji jest audytowana z pełnymi danymi
- Stacja jest widoczna tylko dla Owner, który ją dodał

### US-004: Przeglądanie listy stacji
**Tytuł**: Przeglądanie stacji ładowania

**Opis**: Jako Owner chcę przeglądać listę moich stacji ładowania, aby monitorować moje zasoby.

**Kryteria akceptacji**:
- Owner widzi listę wszystkich swoich stacji
- Lista zawiera podstawowe informacje: Vendor, Model, SerialNumber, status
- Owner może wyszukiwać stacje po nazwie, modelu lub numerze seryjnym
- Lista jest paginowana (np. 20 elementów na stronę)
- Owner może kliknąć na stację, aby zobaczyć szczegóły
- Admin widzi wszystkie stacje w systemie niezależnie od właściciela

### US-005: Edycja stacji ładowania
**Tytuł**: Modyfikacja danych stacji ładowania

**Opis**: Jako Owner chcę móc edytować dane mojej stacji ładowania, aby aktualizować informacje.

**Kryteria akceptacji**:
- Owner może edytować podstawowe dane stacji (Vendor, Model, SerialNumber)
- Owner może dodawać nowe porty do stacji
- Owner może edytować istniejące porty
- Owner może usuwać porty ze stacji
- System waliduje wszystkie wprowadzone dane
- Operacja edycji jest audytowana z informacją o poprzednich i nowych wartościach
- Tylko Owner może edytować swoje stacje, Admin może edytować wszystkie

### US-006: Usuwanie stacji ładowania
**Tytuł**: Usunięcie stacji ładowania z systemu

**Opis**: Jako Owner chcę móc usunąć stację ładowania, która nie jest już potrzebna.

**Kryteria akceptacji**:
- Owner może usunąć stację ładowania
- System sprawdza, czy stacja nie jest przypisana do lokalizacji
- Jeśli stacja jest przypisana do lokalizacji, system wymaga najpierw usunięcia przypisania
- Operacja usunięcia jest audytowana
- Po usunięciu stacji, wszystkie jej porty są również usuwane
- Tylko Owner może usuwać swoje stacje, Admin może usuwać wszystkie

### US-007: Tworzenie nowej lokalizacji
**Tytuł**: Utworzenie lokalizacji fizycznej

**Opis**: Jako Owner chcę utworzyć nową lokalizację fizyczną, aby organizować moje stacje ładowania.

**Kryteria akceptacji**:
- Owner może utworzyć lokalizację podając: Name (string), Address (string), Country (string - ISO 3166-1 alpha-3)
- System generuje unikalny UUID dla lokalizacji
- System waliduje kod kraju zgodnie z ISO 3166-1 alpha-3
- Operacja utworzenia lokalizacji jest audytowana
- Lokalizacja jest widoczna tylko dla Owner, który ją utworzył

### US-008: Przypisanie stacji do lokalizacji
**Tytuł**: Dodanie stacji do lokalizacji

**Opis**: Jako Owner chcę przypisać moje stacje do lokalizacji, aby organizować je według miejsc fizycznych.

**Kryteria akceptacji**:
- Owner może przypisać swoje stacje do lokalizacji
- System umożliwia wyszukiwanie stacji po nazwie/ID podczas przypisywania
- Po przypisaniu stacji do lokalizacji, jej status zmienia się z "w magazynie"
- Owner może przypisać wiele stacji do jednej lokalizacji
- Operacja przypisania jest audytowana
- Tylko Owner może przypisywać swoje stacje, Admin może przypisywać tylko te należące do Ownera

### US-009: Automatyczne generowanie punktów ładowania
**Tytuł**: Generowanie EVSE z portów stacji

**Opis**: Jako system chcę automatycznie generować punkty ładowania (EVSE) z portów stacji przypisanych do lokalizacji.

**Kryteria akceptacji**:
- System automatycznie generuje punkt ładowania dla każdego portu stacji przypisanej do lokalizacji
- Mapowanie jest 1:1 (jeden port = jeden punkt ładowania)
- Każdy punkt ładowania otrzymuje unikalny EvseID zgodny ze specyfikacją Emi3spec
- Punkty ładowania są powiązane z lokalizacją
- Generowanie następuje natychmiast po przypisaniu stacji do lokalizacji
- Operacja generowania jest logowana w systemie audytu

### US-010: Przeglądanie lokalizacji i punktów ładowania
**Tytuł**: Przeglądanie lokalizacji wraz z punktami ładowania

**Opis**: Jako Owner chcę przeglądać moje lokalizacje wraz z przypisanymi stacjami i wygenerowanymi punktami ładowania.

**Kryteria akceptacji**:
- Owner widzi listę swoich lokalizacji
- Dla każdej lokalizacji Owner widzi przypisane stacje
- Dla każdej stacji Owner widzi wygenerowane punkty ładowania (EVSE)
- Lista zawiera podstawowe informacje: nazwa lokalizacji, adres, liczba stacji, liczba punktów ładowania
- Owner może wyszukiwać lokalizacje po nazwie lub adresie
- Admin widzi wszystkie lokalizacje w systemie

### US-011: Edycja lokalizacji
**Tytuł**: Modyfikacja danych lokalizacji

**Opis**: Jako Owner chcę móc edytować dane mojej lokalizacji, aby aktualizować informacje.

**Kryteria akceptacji**:
- Owner może edytować podstawowe dane lokalizacji (Name, Address, Country)
- System waliduje kod kraju zgodnie z ISO 3166-1 alpha-3
- Owner może usuwać przypisania stacji do lokalizacji
- Po usunięciu przypisania, stacja wraca do statusu "w magazynie"
- Operacja edycji jest audytowana z informacją o poprzednich i nowych wartościach
- Tylko Owner może edytować swoje lokalizacje, Admin może edytować wszystkie

### US-012: Usuwanie lokalizacji
**Tytuł**: Usunięcie lokalizacji z systemu

**Opis**: Jako Owner chcę móc usunąć lokalizację, która nie jest już potrzebna.

**Kryteria akceptacji**:
- Owner może usunąć lokalizację
- System sprawdza, czy lokalizacja nie ma przypisanych stacji
- Jeśli lokalizacja ma przypisane stacje, system wymaga najpierw usunięcia wszystkich przypisań
- Po usunięciu lokalizacji, wszystkie wygenerowane punkty ładowania są również usuwane
- Operacja usunięcia jest audytowana
- Tylko Owner może usuwać swoje lokalizacje, Admin może usuwać wszystkie

### US-013: Zarządzanie globalne przez Admin
**Tytuł**: Pełny dostęp Admin do wszystkich zasobów

**Opis**: Jako Admin chcę mieć dostęp do wszystkich stacji i lokalizacji w systemie, aby zarządzać całym systemem.

**Kryteria akceptacji**:
- Admin widzi wszystkie stacje w systemie niezależnie od właściciela
- Admin widzi wszystkie lokalizacje w systemie niezależnie od właściciela
- Admin może edytować, usuwać i przeglądać dowolne zasoby
- Admin może przypisywać stacje do lokalizacji niezależnie od właściciela
- Wszystkie operacje Admin są audytowane z informacją o wykonującym je użytkowniku
- Admin ma dostęp do wszystkich funkcji CRUD dla wszystkich zasobów

### US-014: Wyszukiwanie i filtrowanie zasobów
**Tytuł**: Zaawansowane wyszukiwanie stacji i lokalizacji

**Opis**: Jako użytkownik chcę móc wyszukiwać i filtrować stacje oraz lokalizacje, aby szybko znaleźć potrzebne zasoby.

**Kryteria akceptacji**:
- Użytkownik może wyszukiwać stacje po nazwie, modelu lub numerze seryjnym
- Użytkownik może filtrować stacje według statusu ("w magazynie" vs "przypisana")
- Użytkownik może wyszukiwać lokalizacje po nazwie lub adresie
- Użytkownik może filtrować lokalizacje według kraju
- Wyniki wyszukiwania są paginowane
- Wyszukiwanie działa w czasie rzeczywistym podczas wpisywania
- Owner widzi tylko swoje zasoby, Admin widzi wszystkie

### US-015: Przeglądanie szczegółów punktu ładowania
**Tytuł**: Szczegółowe informacje o punkcie ładowania

**Opis**: Jako użytkownik chcę przeglądać szczegółowe informacje o punkcie ładowania, aby zrozumieć jego specyfikację.

**Kryteria akceptacji**:
- Użytkownik może kliknąć na punkt ładowania, aby zobaczyć szczegóły
- Szczegóły zawierają: EvseID, dane portu (Power, Voltage, Amperage, ConnectorType, ConnectorStandard)
- Szczegóły zawierają informacje o powiązanej stacji i lokalizacji
- Informacje są wyświetlane w czytelnym formacie
- Owner widzi tylko swoje punkty ładowania, Admin widzi wszystkie

### US-016: Audyt operacji systemowych
**Tytuł**: Przeglądanie historii operacji w systemie

**Opis**: Jako Admin chcę przeglądać historię operacji w systemie, aby monitorować aktywność użytkowników.

**Kryteria akceptacji**:
- Admin ma dostęp do logów wszystkich requestów HTTP
- Admin może przeglądać szczegółowy audyt operacji CRUD
- Audyt zawiera informacje o: użytkowniku, operacji, czasie, poprzednich i nowych wartościach
- Logi są filtrowane według użytkownika, typu operacji i zakresu czasowego
- Logi są paginowane dla wydajności
- Dane audytu są przechowywane przez określony okres czasu

## 6. Metryki sukcesu

### 6.1 Metryki adopcji użytkowników

#### Metryka 1: Adopcja stacji ładowania
- **Cel**: 90% użytkowników posiada co najmniej jedną stację ładowania w systemie
- **Okres pomiaru**: 14 dni od rejestracji użytkownika
- **Sposób pomiaru**: Cohortowa analiza użytkowników zarejestrowanych w danym okresie
- **Źródło danych**: Tabela użytkowników + tabela stacji z timestampami
- **Zapytanie**: `COUNT(DISTINCT user_id) WHERE created_chargers >= 1 AND (charger.created_at - user.registered_at) <= 14 dni`
- **Dashboard**: Wykres trendów adopcji w czasie, breakdown per cohort
- **Akcje**: Identyfikacja użytkowników z <1 stacją dla działań retention

#### Metryka 2: Adopcja lokalizacji
- **Cel**: 75% użytkowników posiada co najmniej 3 lokalizacje
- **Okres pomiaru**: 30 dni od rejestracji użytkownika
- **Sposób pomiaru**: Cohortowa analiza użytkowników zarejestrowanych w danym okresie
- **Źródło danych**: Tabela użytkowników + tabela lokalizacji z timestampami
- **Zapytanie**: `COUNT(DISTINCT user_id) WHERE created_locations >= 3 AND (location.created_at - user.registered_at) <= 30 dni`
- **Dashboard**: Wykres postępu, identyfikacja użytkowników z <3 lokalizacjami
- **Akcje**: Działania retention dla użytkowników z niską adopcją lokalizacji

### 6.2 Metryki operacyjne

#### Metryka 3: Czas dodania pierwszej stacji
- **Cel**: Średni czas dodania pierwszej stacji < 5 minut od rejestracji
- **Sposób pomiaru**: Analiza czasowa między rejestracją a dodaniem pierwszej stacji
- **Źródło danych**: Timestampy rejestracji i utworzenia pierwszej stacji
- **Dashboard**: Wykres rozkładu czasów, identyfikacja użytkowników z długim czasem onboardingu

#### Metryka 4: Efektywność wykorzystania stacji
- **Cel**: Monitorowanie stosunku stacji w magazynie do przypisanych
- **Sposób pomiaru**: Stosunek liczby stacji ze statusem "w magazynie" do stacji przypisanych do lokalizacji
- **Źródło danych**: Tabela stacji z polami status
- **Dashboard**: Wykres trendów wykorzystania, alerty przy wysokim odsetku stacji w magazynie

#### Metryka 5: Generowanie punktów ładowania
- **Cel**: Wskaźnik wykorzystania systemu poprzez liczbę wygenerowanych punktów EVSE
- **Sposób pomiaru**: Liczba punktów ładowania wygenerowanych z portów stacji
- **Źródło danych**: Tabela punktów ładowania (EVSE)
- **Dashboard**: Wykres wzrostu liczby punktów ładowania w czasie

### 6.3 Metryki techniczne

#### Metryka 6: Wydajność systemu
- **Cel**: Czas odpowiedzi API < 200ms dla 95% requestów
- **Sposób pomiaru**: Monitoring czasu odpowiedzi wszystkich endpointów API
- **Źródło danych**: Logi serwera web, metryki aplikacji
- **Dashboard**: Wykresy czasu odpowiedzi, identyfikacja wolnych endpointów

#### Metryka 7: Dostępność systemu
- **Cel**: Dostępność systemu > 99.5%
- **Sposób pomiaru**: Monitoring uptime systemu
- **Źródło danych**: Health checks, monitoring infrastruktury
- **Dashboard**: Wykres dostępności, alerty przy awariach

#### Metryka 8: Jakość danych
- **Cel**: < 1% błędów walidacji danych
- **Sposób pomiaru**: Analiza błędów walidacji w logach aplikacji
- **Źródło danych**: Logi aplikacji, błędy walidacji
- **Dashboard**: Wykres błędów walidacji, identyfikacja problematycznych pól

### 6.4 Harmonogram pomiarów
- **Codziennie**: Metryki techniczne (wydajność, dostępność)
- **Tygodniowo**: Metryki operacyjne (czas onboardingu, wykorzystanie stacji)
- **Miesięcznie**: Metryki adopcji (stacje, lokalizacje)
- **Kwartalnie**: Przegląd wszystkich metryk i dostosowanie celów

### 6.5 Kryteria sukcesu MVP
MVP zostanie uznane za udane, jeśli:
1. Metryka 1 (adopcja stacji) osiągnie ≥ 90% w ciągu 14 dni
2. Metryka 2 (adopcja lokalizacji) osiągnie ≥ 75% w ciągu 30 dni
3. Metryka 3 (czas onboardingu) będzie < 5 minut średnio
4. Metryka 6 (wydajność) będzie < 200ms dla 95% requestów
5. Metryka 7 (dostępność) będzie > 99.5%
