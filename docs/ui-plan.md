# Architektura UI dla EV Chargers Management System (ECMS)

## 1. Przegląd struktury UI

EV Chargers Management System jest Single Page Application (SPA) zaprojektowanym jako desktop-first aplikacja z pełną responsywnością. Głównym celem interfejsu jest umożliwienie użytkownikom efektywnego zarządzania stacjami ładowania, lokalizacjami oraz automatycznie generowanymi punktami ładowania (EVSE) poprzez intuicyjny interfejs w Angular 20 z PrimeNG.

### Główne założenia architektury:

- **Desktop-first approach**: Główna nawigacja poprzez sidebar (250px) z możliwością zwijania, z adaptacją na mobile (<992px) jako overlay z hamburger menu
- **Feature-based structure**: Organizacja kodu według funkcjonalności (auth, chargers, locations, dashboard) z lazy loading dla optymalizacji wydajności
- **Angular Signals**: Reaktywne zarządzanie stanem aplikacji bez NgRx w MVP
- **PrimeNG jako biblioteka komponentów**: Wykorzystanie gotowych komponentów z minimalnymi customizations
- **Role-based UI**: Warunkowe renderowanie w zależności od roli użytkownika (Admin vs Owner)
- **Hybrydowe widoki**: Tabele jako główny widok z opcjonalnym widokiem kafelkowym dla elastyczności
- **Centralna obsługa błędów**: Toast dla globalnych komunikatów, inline messages dla walidacji, modals dla błędów krytycznych

## 2. Lista widoków

### 2.1 Autentykacja

#### /auth/login
- **Cel**: Logowanie użytkownika do systemu
- **Kluczowe informacje**:
  - Formularz z polami: email, password
  - Link do rejestracji ("Nie masz konta? Zarejestruj się")
  - Informacja o wymaganiu JWT token
- **Komponenty**:
  - Responsywny card z formularzem (p-card)
  - Input z walidacją emaila (p-inputText)
  - Input password z pokazaniem/ukryciem (p-password)
  - Przycisk "Zaloguj się" z loading state (p-button)
  - Toast messages dla błędów (p-toast)
- **UX**:
  - Wizualny feedback dla błędów walidacji (inline p-message)
  - Loading spinner podczas logowania
  - Automatyczne przekierowanie na /dashboard po udanym logowaniu
  - Zapis JWT w localStorage z automatycznym wylogowaniem po wygaśnięciu (24h)
- **Bezpieczeństwo**:
  - Walidacja client-side przed wysłaniem request
  - Obsługa 401 (błędne dane) z wyświetleniem komunikatu
  - Obsługa błędu sieci/połączenia

#### /auth/register
- **Cel**: Rejestracja nowego użytkownika
- **Kluczowe informacje**:
  - Formularz z polami: email, password, confirm password
  - Informacja o automatycznej roli Owner
  - Link do logowania ("Masz już konto? Zaloguj się")
- **Komponenty**:
  - Responsywny card z formularzem (p-card)
  - Input z walidacją emaila (p-inputText)
  - Input password z pokazaniem/ukryciem i metryką siły (p-password)
  - Input confirm password z walidacją zgodności (p-inputText)
  - Przycisk "Zarejestruj się" z loading state (p-button)
  - Toast messages dla błędów (p-toast)
- **UX**:
  - Client-side validation z natychmiastowym feedbackiem
  - Wskaźnik siły hasła (min. 8 znaków)
  - Walidacja zgodności hasła
  - Automatyczne logowanie i przekierowanie po rejestracji
- **Bezpieczeństwo**:
  - Walidacja formatu emaila
  - Walidacja minimalnej długości hasła (8 znaków)
  - Obsługa 409 (email już istnieje)
  - Obsługa błędów walidacji z backend (400)

### 2.2 Dashboard

#### /dashboard
- **Cel**: Główny punkt wejścia aplikacji po zalogowaniu, prezentujący przegląd zasobów
- **Kluczowe informacje**:
  - Zamockowane statystyki:
    - Liczba stacji ogółem (total chargers)
    - Liczba stacji w magazynie (warehouse chargers)
    - Liczba lokalizacji (total locations)
    - Liczba punktów EVSE (total EVSE)
  - Ostatnie aktywności w systemie
  - Szybkie akcje (kafelki do przejścia do sekcji)
- **Komponenty**:
  - Stat cards z ikonami i liczbami (custom cards z PrimeNG grid)
  - Panel ostatnich aktywności (p-table lub custom list)
  - Kafelki szybkich akcji (p-button jako tiles)
- **UX**:
  - Responsywny grid: 4 kolumny desktop, 2 tablet, 1 mobile
  - Hover effects na kafelkach akcji
  - Skeleton loader podczas ładowania danych
  - Click na kafelki → nawigacja do odpowiednich sekcji
- **Bezpieczeństwo**:
  - Wymaga autentykacji (AuthGuard)
  - Role-aware dane (Admin widzi wszystko, Owner tylko swoje)

### 2.3 Stacje ładowania

#### /chargers
- **Cel**: Lista wszystkich stacji ładowania z możliwością przeszukiwania i filtrowania
- **Kluczowe informacje**:
  - Lista stacji z kolumnami: Vendor, Model, Serial Number, Status, Connectors Count
  - Liczba connectorów jako badge
  - Status: badge "W magazynie" / "Przypisana"
  - Action buttons: View, Edit, Delete (inline lub w dropdown menu)
- **Komponenty**:
  - Panel filtrów (p-inputText dla search, p-dropdown dla status, p-button "Wyczyść")
  - Toggle view (tabela/kafelki) (p-toggleButton)
  - Tabela z paginacją (p-table z paginator lub p-dataView)
  - Przycisk "Dodaj stację" (p-button)
  - Toast dla operacji delete
  - Confirm dialog dla delete (p-confirmDialog)
- **Komponenty filtrowania**:
  - Search input z debounce 300ms (vendor, model, serial_number)
  - Dropdown status ("wszystkie", "w magazynie", "przypisane")
  - Dropdown location (jeśli Admin)
  - Badge z liczbą aktywnych filtrów
  - Przycisk "Wyczyść wszystkie"
- **UX**:
  - Real-time search podczas wpisywania (debounce)
  - Paginacja: 20/50/100 items per page
  - Kliknięcie w wiersz → /chargers/:id
  - Hybrydowy widok (tabela default, toggle do kafelków)
  - Sortowanie po kolumnach
  - Empty state gdy brak stacji
  - Loading state podczas fetch
- **Bezpieczeństwo**:
  - Automatyczne filtrowanie według roli (Owner widzi tylko swoje)
  - Role-aware delete (weryfikacja backend)
  - Obsługa błędu przy usuwaniu stacji przypisanej do lokalizacji (400)

#### /chargers/new
- **Cel**: Tworzenie nowej stacji ładowania z connectorami
- **Kluczowe informacje**:
  - Podstawowe dane: Vendor, Model, Serial Number
  - Dynamiczna lista connectorów (min. 1): ConnectorID, Power, Voltage, Amperage, ConnectorType, ConnectorStandard
- **Komponenty**:
  - Reactive Form z FormArray dla connectorów
  - Input dla podstawowych danych (p-inputText)
  - Fieldset dla każdego connectora (p-fieldset)
  - Dropdown dla ConnectorType (CCS, Type2, Chademo)
  - Dropdown dla ConnectorStandard (AC_1P, AC_3P, DC)
  - Numeryczne inputy z walidacją (p-inputNumber)
  - Przycisk "Dodaj connector" (p-button)
  - Przycisk "Usuń connector" z potwierdzeniem (p-button)
  - Przycisk "Zapisz" z loading state (p-button)
  - Przycisk "Anuluj" → powrót do listy (p-button)
  - CanDeactivate guard z dirty check
  - Toast dla success
- **UX**:
  - Client-side validation z natychmiastowym feedbackiem
  - Badge "Istniejący" / "Nowy" dla connectorów
  - Visual indicator dla wymaganych pól (asterisk)
  - Wskaźnik liczby connectorów
  - Skeleton loader podczas submit
- **Bezpieczeństwo**:
  - Walidacja wszystkich pól przed submit
  - Walidacja unikalności serial_number
  - Obsługa 409 (serial number już istnieje)
  - Walidacja positive numbers dla power, voltage, amperage

#### /chargers/:id
- **Cel**: Szczegóły pojedynczej stacji z możliwością edycji i zarządzania connectorami
- **Kluczowe informacje**:
  - Breadcrumbs: "Stacje > [Vendor] [Model] [SerialNumber]"
  - Podstawowe dane stacji (read-only z możliwością przejścia do edycji)
  - Lista wszystkich connectorów z szczegółami
  - Status i przypisanie do lokalizacji
- **Komponenty**:
  - Breadcrumbs (p-breadcrumb)
  - Header z tytułem i akcjami (Edit, Delete)
  - Panel informacji podstawowych (p-panel)
  - Tabela/l lista connectorów (p-table lub p-dataList)
  - Przycisk "Dodaj connector"
  - Przycisk "Edytuj connector" dla każdego
  - Przycisk "Usuń connector" z potwierdzeniem
  - Detail view dla connectorów z rozszerzalnymi wierszami (expandable rows)
  - Modal do edycji connectorów
- **UX**:
  - Rozszerzane wiersze dla szczegółów connectora
  - Inline edycja connectorów
  - Badge statusowe (Istniejący/Nowy)
  - Loading state podczas fetch i update
  - Success toast po każdej operacji
- **Bezpieczeństwo**:
  - Version check przy edycji (optimistic locking)
  - Obsługa 409 Conflict (version mismatch)
  - Modal przy konflikcie z opcją odświeżenia danych
  - Walidacja przed submit zmian

#### /chargers/:id/edit
- **Cel**: Pełna edycja stacji z connectorami
- **Kluczowe informacje**:
  - Formularz z podstawowymi danymi
  - Dynamiczna lista connectorów z edycją
  - Version field (hidden w UI, używane przy submit)
- **Komponenty**:
  - Podobne do /chargers/new, ale z wstępnie wypełnionymi danymi
  - Badge dla istniejących connectorów
  - Optional delete dla connectorów
  - Modal confirm dla usunięcia istniejącego connectora
  - Submit z wersją
- **UX**:
  - Visual indicators dla connectorów (kolor lub badge)
  - Dirty check w CanDeactivate guard
  - Loading state podczas optimistic locking check
  - Success/error feedback
- **Bezpieczeństwo**:
  - Optimistic locking (409 Conflict)
  - Conflict modal z opcją refresh
  - Walidacja client-side i backend

### 2.4 Lokalizacje

#### /locations
- **Cel**: Lista wszystkich lokalizacji z możliwością przeszukiwania i filtrowania
- **Kluczowe informacje**:
  - Lista lokalizacji z kolumnami: Name, Address, Country, Chargers Count, EVSE Count
  - Liczba przypisanych stacji jako badge
  - Liczba wygenerowanych punktów EVSE
  - Action buttons: View, Edit, Delete
- **Komponenty**:
  - Panel filtrów (p-inputText dla search name/address, p-dropdown dla country, p-button "Wyczyść")
  - Tabela z paginacją (p-table)
  - Przycisk "Dodaj lokalizację" (p-button)
  - Toast dla operacji delete
  - Confirm dialog dla delete
- **UX**:
  - Real-time search podczas wpisywania (debounce)
  - Paginacja: 20/50/100 items per page
  - Kliknięcie w wiersz → /locations/:id
  - Hybrydowy widok (tabela default)
  - Sortowanie po kolumnach
  - Empty state gdy brak lokalizacji
- **Bezpieczeństwo**:
  - Automatyczne filtrowanie według roli (Owner widzi tylko swoje)
  - Role-aware delete
  - Obsługa błędu przy usuwaniu lokalizacji z przypisanymi stacjami (400)

#### /locations/new
- **Cel**: Tworzenie nowej lokalizacji fizycznej
- **Kluczowe informacje**:
  - Pola: Name, Address, Country (ISO 3166-1 alpha-3)
- **Komponenty**:
  - Reactive Form
  - Input dla Name (p-inputText)
  - Textarea dla Address (p-inputTextarea)
  - Input dla Country z walidacją ISO 3166-1 alpha-3 (p-inputText z maską 3 uppercase letters)
  - Przycisk "Zapisz" z loading state
  - Przycisk "Anuluj"
  - CanDeactivate guard z dirty check
  - Toast dla success
- **UX**:
  - Client-side validation z natychmiastowym feedbackiem
  - Regex validation dla country code (3 uppercase letters)
  - Visual indicator dla wymaganych pól
  - Auto-redirect do detail view po utworzeniu
- **Bezpieczeństwo**:
  - Walidacja formatu country code (ISO 3166-1 alpha-3)
  - Walidacja wymaganych pól
  - Backend validation jako ostateczna bariera

#### /locations/:id
- **Cel**: Szczegóły lokalizacji z zakładkami dla stacji i EVSE
- **Kluczowe informacje**:
  - Breadcrumbs: "Lokalizacje > [Name]"
  - Zakładki: Informacje, Stacje, EVSE
  - Przycisk "Przypisz stację" (tylko w zakładce Stacje)
- **Komponenty**:
  - Breadcrumbs (p-breadcrumb)
  - TabView (p-tabView) z trzema zakładkami
  - Przycisk "Edytuj" i "Usuń" w header
  - Panel informacji (p-panel)
  - Tabela stacji z expandable rows (p-table)
  - Tabela EVSE (p-table)
  - Modal "Przypisz stację" z typeahead (p-dialog, p-autoComplete)
  - Toast dla operacji assign/delete
- **Zakładka "Informacje"**:
  - Nazwa, adres, kraj (edytowalne inline lub w osobnym widoku)
  - Version field (hidden, dla optimistic locking)
- **Zakładka "Stacje"**:
  - Lista przypisanych stacji z expandable rows dla connectorów
  - Badge z liczbą connectorów
  - Przycisk "Przypisz stację" (tylko Owner/Admin)
  - Modal typeahead search dla wyboru stacji
  - Przycisk "Odepnij stację" dla każdej
  - Empty state gdy brak stacji
- **Zakładka "EVSE"**:
  - Lista wygenerowanych punktów EVSE
  - Kolumny: EvseID, Connector details, Created At
  - Mapowanie 1:1 connector → EVSE
  - Read-only (generowane automatycznie)
  - Empty state gdy brak EVSE
- **UX**:
  - Skeleton loader podczas ładowania
  - Rozszerzane wiersze dla szczegółów stacji/connectorów
  - Success toast po przypisaniu/odepnieniu stacji
  - Auto-refresh EVSE po przypisaniu stacji
  - Loading state podczas generowania EVSE
- **Bezpieczeństwo**:
  - Role-aware przypisanie stacji (Owner tylko swoje stacje)
  - Walidacja przed przypisaniem (czy stacja już przypisana)
  - Obsługa błędów przy przypisaniu
  - Version check przy edycji lokalizacji

#### /locations/:id/assign-charger (Modal)
- **Cel**: Przypisanie stacji do lokalizacji z typeahead search
- **Kluczowe informacje**:
  - Wyszukiwanie stacji po vendor, model, serial_number
  - Podgląd wybranej stacji przed przypisaniem
  - Lista connectorów wybranej stacji
- **Komponenty**:
  - p-dialog jako modal
  - p-autoComplete dla wyszukiwania stacji (tylko stacje "w magazynie")
  - Panel podglądu stacji (p-panel)
  - Lista connectorów (p-dataList lub custom list)
  - Przycisk "Przypisz" z loading state
  - Przycisk "Anuluj"
- **UX**:
  - Real-time search podczas wpisywania
  - Highlight wybranej stacji
  - Podgląd connectorów przed przypisaniem
  - Success feedback po przypisaniu
  - Auto-close modal i refresh widoku
- **Bezpieczeństwo**:
  - Filtrowanie tylko stacji "w magazynie"
  - Walidacja czy stacja należy do użytkownika
  - Obsługa 400 (stacja już przypisana)
  - Auto-generacja EVSE po przypisaniu

### 2.5 Profil (opcjonalny w MVP)

#### /profile
- **Cel**: Zarządzanie danymi użytkownika
- **Kluczowe informacje**:
  - Email (read-only w MVP)
  - Rola (read-only, wyświetlana jako badge)
  - Data rejestracji
  - Przycisk "Wyloguj się"
- **Komponenty**:
  - Panel z informacjami użytkownika (p-panel)
  - Badge z rolą (Admin/Owner)
  - Przycisk "Wyloguj się" (p-button)
  - Confirm dialog dla wylogowania
- **UX**:
  - Read-only w MVP (brak edycji)
  - Czytelny layout z danymi
- **Bezpieczeństwo**:
  - Wylogowanie usuwa JWT z localStorage
  - Przekierowanie na /auth/login

## 3. Mapa podróży użytkownika

### 3.1 Nowy użytkownik - Onboarding (US-003, US-007, US-008, US-009)

**Cel**: Przewodnik nowego użytkownika przez pierwsze kroki w systemie

**Przepływ**:
1. Rejestracja (/auth/register) → Automatyczne logowanie i przekierowanie
2. Dashboard (/dashboard) → Przegląd zamockowanych statystyk, orientacja w systemie
3. Dodanie pierwszej stacji (/chargers/new):
   - Formularz podstawowych danych (vendor, model, serial_number)
   - Dodanie minimum 1 connectora z pełnymi szczegółami
   - Zapisz → Success toast → Przekierowanie na /chargers
4. Lista stacji (/chargers) → Weryfikacja dodanej stacji
5. Szczegóły stacji (/chargers/:id) → Przegląd danych i connectorów
6. Dodanie pierwszej lokalizacji (/locations/new):
   - Formularz nazwy, adresu, kodu kraju
   - Zapisz → Success toast → Przekierowanie na /locations/:id
7. Przypisanie stacji do lokalizacji (/locations/:id):
   - Zakładka "Stacje" → "Przypisz stację"
   - Modal typeahead → Wybór stacji z podglądem
   - Potwierdzenie → Automatyczne generowanie EVSE
8. Weryfikacja wygenerowanych EVSE (/locations/:id, zakładka "EVSE"):
   - Lista punktów ładowania z EvseID
   - Weryfikacja mapowania 1:1 connector → EVSE

### 3.2 Codzienne zarządzanie zasobami (US-004, US-005, US-010, US-014)

**Cel**: Efektywne zarządzanie stacjami i lokalizacjami

**Przepływy**:

**Wyszukiwanie i filtrowanie stacji**:
- Dashboard → Stacje (/chargers)
- Wprowadzenie search term w panel filtrów (debounce 300ms)
- Filtrowanie po statusie (dropdown)
- Przegląd wyników w czasie rzeczywistym
- Kliknięcie w wiersz → Szczegóły stacji

**Edycja stacji**:
- Lista stacji → Kliknięcie w wiersz → /chargers/:id
- Przycisk "Edytuj" → /chargers/:id/edit
- Modyfikacja danych i connectorów (dodawanie/modyfikowanie/usuwanie)
- Zapisz → Optimistic locking check → Success/Conflict handling
- Powrót do szczegółów z zaktualizowanymi danymi

**Przegląd lokalizacji z EVSE**:
- Dashboard → Lokalizacje (/locations)
- Wyszukiwanie po nazwie/adresie
- Filtrowanie po kraju
- Kliknięcie w wiersz → /locations/:id
- Przegląd zakładek: Informacje, Stacje, EVSE
- Rozszerzenie wierszy stacji dla connectorów
- Weryfikacja wygenerowanych punktów EVSE

### 3.3 Przypisywanie/detach stacji (US-008)

**Cel**: Organizacja stacji według lokalizacji fizycznych

**Przepływ przypisywania**:
1. Lokalizacje → Szczegóły lokalizacji (/locations/:id)
2. Zakładka "Stacje" → "Przypisz stację"
3. Modal opens:
   - Typeahead search dla stacji "w magazynie"
   - Podgląd wybranej stacji
   - Lista connectorów
4. Potwierdzenie → PUT /locations/:id/assign-charger
5. Auto-generacja EVSE
6. Auto-refresh widoku z nową stacją i EVSE
7. Success toast

**Przepływ detach**:
1. Lokalizacje → Szczegóły → Zakładka "Stacje"
2. Przycisk "Odepnij stację" przy wybranej stacji
3. Confirm dialog
4. DELETE /locations/:id/chargers/:charger_id
5. Auto-delete powiązanych EVSE
6. Stacja wraca do statusu "w magazynie"
7. Success toast

### 3.4 Usuwanie zasobów (US-006, US-012)

**Cel**: Bezpieczne usuwanie stacji i lokalizacji

**Przepływ usuwania stacji**:
1. Lista stacji → Kliknięcie delete (inline lub dropdown)
2. Confirm dialog z informacją o konsekwencjach
3. Sprawdzenie czy stacja jest przypisana:
   - Jeśli nie → Soft delete (DELETE /chargers/:id)
   - Jeśli tak → Błąd 400, komunikat o konieczności najpierw usunięcia przypisania
4. Success toast
5. Refresh listy

**Przepływ usuwania lokalizacji**:
1. Lista lokalizacji → Kliknięcie delete
2. Confirm dialog z informacją o konsekwencjach
3. Sprawdzenie czy lokalizacja ma przypisane stacje:
   - Jeśli nie → Soft delete (DELETE /locations/:id)
   - Jeśli tak → Błąd 400, komunikat o konieczności najpierw odepnięcia stacji
4. Auto-delete powiązanych EVSE
5. Success toast
6. Refresh listy

### 3.5 Zarządzanie globalne przez Admin (US-013)

**Cel**: Admin ma pełny dostęp do wszystkich zasobów

**Różnice w interfejsie dla Admina**:
- Wszystkie widoki list pokazują wszystkie zasoby (bez filtrowania po owner_id)
- Admin może edytować/usuwać zasoby wszystkich użytkowników
- W widokach szczegółów widoczny owner_id zasobu
- Badge "Admin" w top bar (opcjonalnie)
- Wszystkie operacje Admin są audytowane z user_id w logach

**Przepływ dla Admina**:
1. Login jako Admin
2. Dashboard → Widok statystyk całego systemu
3. Stacje → Widok wszystkich stacji (nie tylko swoich)
4. Lokalizacje → Widok wszystkich lokalizacji
5. Operacje CRUD na zasobach wszystkich użytkowników
6. Wszystkie operacje logowane w audyt trail

### 3.6 Wyszukiwanie zaawansowane (US-014)

**Cel**: Szybkie znajdowanie zasobów w dużych zbiorach

**Przepływ wyszukiwania stacji**:
- Lista stacji → Panel filtrów
- Search input (real-time, debounce 300ms) → Vendor, Model, Serial Number
- Status dropdown → "wszystkie", "w magazynie", "przypisane"
- Location dropdown (Admin only)
- Badge z liczbą aktywnych filtrów
- Kliknięcie "Wyczyść" → Reset wszystkich filtrów

**Przepływ wyszukiwania lokalizacji**:
- Lista lokalizacji → Panel filtrów
- Search input (real-time, debounce 300ms) → Name, Address
- Country dropdown → Filtr po kodzie kraju
- Badge z liczbą aktywnych filtrów

### 3.7 Obsługa błędów i konfliktów

**Optimistic Locking Conflict (409)**:
- Podczas PUT /chargers/:id lub /locations/:id
- Conflict modal z informacją o konflikcie wersji
- Opcje: "Odśwież dane" (GET latest version) lub "Anuluj"
- Visual indicator w UI o konieczności odświeżenia

**Błędy walidacji (400)**:
- Client-side validation z inline messages
- Backend validation errors mapowane na form controls
- Toast dla ogólnych błędów
- p-message dla błędów przy konkretnych polach

**Błędy autoryzacji (401, 403)**:
- 401 → Automatyczne wylogowanie + redirect na /auth/login
- 403 → Toast "Brak uprawnień"
- 404 → Toast lub dedicated 404 page

## 4. Układ i struktura nawigacji

### 4.1 Główny layout

**Struktura aplikacji**:
```
┌─────────────────────────────────────────────────┐
│ Top Bar: Logo | Breadcrumbs | User Menu | Admin Badge │
├─────────────────────────────────────────────────┤
│            │                                    │
│   Sidebar  │          Main Content             │
│            │                                    │
│   (250px)  │      (Resizable Area)             │
│            │                                    │
└─────────────────────────────────────────────────┘
```

### 4.2 Sidebar Navigation

**Desktop (≥992px)**:
- Persistent sidebar 250px szerokości
- Logo w headerze
- Menu items: Dashboard, Stacje, Lokalizacje
- Przycisk zwijania (chevron left/right)
- Stan (expanded/collapsed) zapisywany w localStorage
- Smooth transition podczas zwijania

**Mobile (<992px)**:
- Hamburger menu w top bar
- Sidebar jako overlay z backdrop
- Slide-in animation (from left)
- Touch swipe do zamknięcia
- Auto-close po kliknięciu w link

**Menu items**:
- **Dashboard** (/dashboard) - Ikona home, zawsze aktywny po zalogowaniu
- **Stacje** (/chargers) - Ikona plug/battery, dostęp do stacji
- **Lokalizacje** (/locations) - Ikona map-marker, dostęp do lokalizacji

**Brak**:
- Osobnej sekcji EVSE (dostępne tylko w kontekście lokalizacji)
- Liczników zasobów w menu (decyzja UX)
- Sub-menu items

### 4.3 Breadcrumbs

**Implementacja**:
- Routing configuration z route.data dla breadcrumb labels
- BreadcrumbService subskrybuje router events
- PrimeNG p-breadcrumb component
- Przykład: "Stacje > Tesla Supercharger V3"
- Kliknięcie w breadcrumb → Nawigacja do odpowiedniej ścieżki

**Scenariusze**:
- Root level (Dashboard) → Tylko "Dashboard"
- Nested (Stacja) → "Stacje > Vendor Model"
- Deep nested (Lokalizacja → Stacje) → "Lokalizacje > Name"

### 4.4 Top Bar

**Elementy**:
- Logo aplikacji (lewym górnym rogu)
- Breadcrumbs (centrum, dynamiczne)
- User menu (prawym górnym rogu):
  - Ikona user
  - Email użytkownika lub akronim
  - Dropdown: Profil, Wyloguj się
- Progress bar (na samej górze, podczas navigation)

**Responsywność**:
- Mobile: Logo + Hamburger menu (po lewej), User menu (po prawej)
- Tablet: Pełna top bar z responsive breadcrumbs
- Desktop: Pełna top bar z wszystkimi elementami

### 4.5 Routing Structure

```
/auth
  /login           → LoginPageComponent
  /register        → RegisterPageComponent

/dashboard         → DashboardComponent (AuthGuard)

/chargers          → ChargersListComponent (AuthGuard)
  /new             → ChargerFormComponent (AuthGuard)
  /:id              → ChargerDetailComponent (AuthGuard)
    /edit          → ChargerEditComponent (AuthGuard, CanDeactivateGuard)

/locations         → LocationsListComponent (AuthGuard)
  /new             → LocationFormComponent (AuthGuard)
  /:id              → LocationDetailComponent (AuthGuard)
    /edit          → LocationEditComponent (AuthGuard, CanDeactivateGuard)

/profile           → ProfileComponent (AuthGuard, optional in MVP)

NotFoundComponent  → 404 page
```

**Route Guards**:
- **AuthGuard**: Sprawdza JWT token, redirect na /auth/login jeśli nie autentykowany
- **CanDeactivateGuard**: Sprawdza dirty forms, confirm przed opuszczeniem
- Role-based guards nie są potrzebne (backend enforces authorization)

### 4.6 Deep Linking

**Wszystkie widoki wspierają deep linking**:
- Każda akcja użytkownika (click na wiersz) przekierowuje do szczegółów
- URL zawiera ID zasobu (/chargers/:id, /locations/:id)
- Refresh strony zachowuje stan
- Shareable URLs (dla Admin lub Owner zasobów)

## 5. Kluczowe komponenty

### 5.1 Layout Components

**AppLayoutComponent**:
- Główny wrapper dla całej aplikacji (po zalogowaniu)
- Zawiera: Sidebar, TopBar, MainContent area
- Zarządza responsive breakpoint 992px
- Obsługuje state sidebar (expanded/collapsed)
- Przechowuje JWT w localStorage

**SidebarComponent**:
- Menu nawigacyjne z linkami
- Obsługa collapse/expand
- Responsive behavior (persistent desktop, overlay mobile)
- Active route highlighting
- Smooth transitions

**TopBarComponent**:
- Header z logo, breadcrumbs, user menu
- Breadcrumbs z dynamic labels
- User dropdown z logout
- Progress bar dla navigation
- Mobile hamburger toggle

### 5.2 Auth Components

**LoginFormComponent**:
- Reactive form z email, password
- Client-side validation
- Error handling z toast i inline messages
- Loading state podczas submit
- Auto-redirect po sukcesie

**RegisterFormComponent**:
- Reactive form z email, password, confirm password
- Password strength indicator
- Email format validation
- Confirm password match validation
- Success flow z auto-login

### 5.3 Charger Components

**ChargersListComponent**:
- Panel filtrów z search i dropdowns
- Toggle view (table/tiles)
- p-table z paginacją
- Inline actions (view, edit, delete)
- Empty state
- Loading skeleton
- Toast messages

**ChargerFormComponent** (create/edit):
- Reactive Form z FormArray dla connectorów
- Dynamiczne dodawanie/usuwanie connectorów
- Validation dla wszystkich pól
- Badge "Istniejący" / "Nowy"
- CanDeactivate guard
- Version field dla optimistic locking

**ChargerDetailComponent**:
- Header z akcjami (edit, delete)
- Lista connectorów z expandable rows
- Badge statusowe
- Version field hidden
- Conflict modal

### 5.4 Location Components

**LocationsListComponent**:
- Panel filtrów z search i country dropdown
- p-table z paginacją
- Inline actions
- Empty state
- Loading skeleton

**LocationFormComponent** (create/edit):
- Reactive Form z Name, Address, Country
- ISO 3166-1 alpha-3 validation
- CanDeactivate guard
- Version field hidden

**LocationDetailComponent**:
- TabView z zakładkami: Informacje, Stacje, EVSE
- Panel informacji
- Przypisywanie stacji (modal)
- Lista stacji z expandable rows
- Lista EVSE (read-only)
- Version field hidden

**AssignChargerModalComponent**:
- Typeahead search dla stacji "w magazynie"
- Podgląd wybranej stacji
- Lista connectorów
- Submit z loading state
- Auto-close i refresh

### 5.5 Shared Components

**ConfirmDialogComponent** (p-confirmDialog):
- Potwierdzenie destrukcyjnych akcji
- Customizable title i message
- Przycisk "Tak" / "Nie"
- Keyboard navigation (Enter, Escape)

**ToastComponent** (p-toast):
- Globalne komunikaty o sukcesie/błędzie
- Position top-right default
- Auto-dismiss po 3 sekundach
- Stacking dla wielu komunikatów

**LoadingSpinnerComponent**:
- Local spinners dla button loading states
- Skeleton screens dla empty state
- Top progress bar dla navigation
- Reusable loading overlay

**EmptyStateComponent**:
- Komunikat gdy brak danych
- Obrazek/ikonka
- Call-to-action (np. "Dodaj pierwszą stację")
- Conditional rendering na podstawie state

**ErrorMessageComponent** (p-message):
- Inline error messages dla form controls
- Walidacja client-side
- ARIA attributes dla accessibility
- Positioning pod polem

**VersionConflictModalComponent**:
- Obsługa 409 Conflict dla optimistic locking
- Informacja o konflikcie wersji
- Opcje: "Odśwież dane" lub "Anuluj"
- Auto-refresh dane przy wyborze odświeżenia

### 5.6 Services

**AuthService**:
- Login/Register logic
- JWT storage w localStorage
- Token expiration check
- Logout functionality
- Signals dla isAuthenticated, currentUser

**ChargersService**:
- CRUD dla stacji
- List z paginacją i filtrami
- Search functionality
- Signals dla chargers, loading, error states
- Optimistic locking handling

**LocationsService**:
- CRUD dla lokalizacji
- List z paginacją i filtrami
- Assign/detach chargers
- Fetch EVSE for location
- Signals dla locations, loading, error states

**BreadcrumbService**:
- Dynamic breadcrumb labels from route.data
- Subscription to router events
- Breadcrumb path building
- Signal dla breadcrumbs

**ErrorInterceptor**:
- Centralizacja obsługi błędów HTTP
- 401 → Logout + redirect
- 403 → Toast
- 409 → Version conflict modal
- 5xx → Toast z ogólnym komunikatem

**AuthInterceptor**:
- Dodaje JWT token do HTTP headers
- Interceptuje requesty
- Skonfigurowany per request

### 5.7 Guards

**AuthGuard**:
- Sprawdza JWT token w localStorage
- Redirect na /auth/login jeśli nie autentykowany
- Wszystkie chronione trasy

**CanDeactivateGuard**:
- Sprawdza dirty state forms
- Confirm dialog przed opuszczeniem
- Używane w /chargers/new, /chargers/:id/edit, /locations/new, /locations/:id/edit

### 5.8 Directives & Pipes

**DebounceDirective** (dla search inputs):
- Debounce 300ms dla zmniejszenia API calls
- Używane w filtrze search

**RoleDirective** (NgIf-based):
- Conditional rendering według roli
- Używa parsed JWT role
- Użyte w wielu widokach

**ConnectorTypePipe**:
- Formatowanie typów connectorów
- Użyte w display connector lists

**ConnectorStandardPipe**:
- Formatowanie standardów connectorów
- Użyte w display connector details

### 5.9 Models & Interfaces

**Charger** (domain model):
- Podstawowe pola: vendor, model, serial_number
- Status: warehouse/assigned
- Location_id (nullable)
- Version dla optimistic locking
- Connectors array

**Location** (domain model):
- Name, Address, Country
- Version dla optimistic locking
- Chargers array
- EVSE array

**Connector** (domain model):
- ConnectorID, Power, Voltage, Amperage
- ConnectorType, ConnectorStandard
- ID dla edit (nullable w create)

**EVSE** (domain model):
- EvseID (Emi3spec format)
- Connector relationship
- Created_at

**AuthState** (JWT payload):
- User ID
- Email
- Role (admin/owner)

## 6. UX/UI Best Practices

### 6.1 Response Times
- API response target: <200ms dla 95% requestów
- Skeleton loaders dla waiting states
- Progressive loading danych (lazy loading modules)
- Debounce dla search (300ms)

### 6.2 Feedback
- Toast messages dla sukcesów/błędów
- Inline validation messages
- Loading indicators na buttonach
- Skeleton screens zamiast blank screens
- Progress bar dla navigation
- Success sounds (optional)

### 6.3 Error Handling
- Graceful degradation
- Retry logic dla network errors (future)
- User-friendly error messages
- Conflict resolution UI dla optimistic locking
- 404 page dla nieistniejących zasobów

### 6.4 Accessibility
- ARIA labels dla wszystkich interakcji
- Keyboard navigation
- Focus management w modalach
- Color contrast WCAG AA
- Screen reader support
- Semantic HTML

### 6.5 Mobile Optimization
- Touch-friendly controls (min 44px)
- Swipe gestures dla sidebar
- Responsive tables (priorytet kolumn)
- Stacked forms na mobile
- Bottom sheet zamiast modals (future consideration)

## 7. Integration Points

### 7.1 API Integration
- HttpClient z interceptors
- Typed responses (TypeScript interfaces)
- Error handling w ErrorInterceptor
- Optimistic locking w PUT requests
- Version field management

### 7.2 State Management
- Angular Signals dla reactive state
- Services per feature
- Computed signals dla derived state
- No NgRx w MVP (mogąc migrate w przyszłości)

### 7.3 Styling
- PrimeNG theme (Lara Light Blue/Dark)
- Custom SCSS variables
- Responsive breakpoints
- Consistent spacing system

## 8. Mapping User Stories to UI

### US-001: Rejestracja → /auth/register
- Formularz email + password
- Auto-redirect na dashboard po sukcesie

### US-002: Logowanie → /auth/login
- Formularz email + password
- JWT storage + auto-redirect

### US-003: Dodanie pierwszej stacji → /chargers/new
- Formularz z FormArray dla connectorów
- Success toast + redirect do listy

### US-004: Przeglądanie listy → /chargers
- Panel filtrów + tabela
- Paginacja i search

### US-005: Edycja stacji → /chargers/:id/edit
- Formularz z version field
- Optimistic locking handling

### US-006: Usuwanie stacji → /chargers
- Confirm dialog
- Błąd 400 jeśli przypisana

### US-007: Tworzenie lokalizacji → /locations/new
- Formularz Name, Address, Country
- ISO validation

### US-008: Przypisanie stacji → /locations/:id (modal)
- Typeahead search
- Auto-generacja EVSE

### US-009: Generowanie EVSE → Automatyczne
- Po przypisaniu stacji
- Widoczne w zakładce "EVSE"

### US-010: Przegląd lokalizacji → /locations/:id
- TabView z zakładkami
- Rozszerzane wiersze dla stacji

### US-011: Edycja lokalizacji → /locations/:id
- Edit inline lub osobny widok
- Version field

### US-012: Usuwanie lokalizacji → /locations
- Confirm dialog
- Błąd 400 jeśli ma stacje

### US-013: Zarządzanie globalne → Wszystkie widoki
- Role-aware rendering
- Admin badge w UI
- Wszystkie operacje audytowane

### US-014: Wyszukiwanie → Panel filtrów w listach
- Search + debounce
- Filters + badge z liczbą aktywnych

### US-015: Szczegóły EVSE → /locations/:id (tab "EVSE")
- Tabela z EvseID i detalami
- Read-only display

### US-016: Audyt → Poza MVP UI
- Backend logowanie
- Future: dedicated audit logs view dla Admin

## 9. Technical Implementation Notes

### 9.1 File Structure (Feature-based)
```
frontend/src/
  app/
    features/
      auth/
        components/
          login-form/
          register-form/
        services/
          auth.service.ts
        guards/
          auth.guard.ts
        models/
          auth.models.ts
      chargers/
        components/
          chargers-list/
          charger-form/
          charger-detail/
        services/
          chargers.service.ts
        guards/
          can-deactivate.guard.ts
        models/
          charger.models.ts
      locations/
        components/
          locations-list/
          location-form/
          location-detail/
          assign-charger-modal/
        services/
          locations.service.ts
        models/
          location.models.ts
      dashboard/
        components/
          dashboard/
        services/
          dashboard.service.ts
    core/
      interceptors/
        auth.interceptor.ts
        error.interceptor.ts
      guards/
        can-deactivate.guard.ts
      services/
        breadcrumb.service.ts
    shared/
      components/
        confirm-dialog/
        empty-state/
        loading-spinner/
        version-conflict-modal/
      directives/
        debounce.directive.ts
        role.directive.ts
      pipes/
        connector-type.pipe.ts
      models/
        pagination.models.ts
    layout/
      components/
        app-layout/
        sidebar/
        topbar/
    app.config.ts
    app.routes.ts
```

### 9.2 Lazy Loading Modules
- Każdy feature lazy loaded
- Routes configuration z loadComponent
- Redukcja initial bundle size

### 9.3 Testing Strategy (future)
- Unit tests dla services
- Component tests dla forms
- Integration tests dla workflows
- E2E tests dla critical paths

---

## Podsumowanie

Architektura UI dla ECMS została zaprojektowana z myślą o intuicyjnym doświadczeniu użytkownika, efektywnym zarządzaniu zasobami oraz wsparciem dla różnych ról (Owner i Admin). Główne filary architektury to desktop-first approach, feature-based organization, Angular Signals dla state management, oraz PrimeNG jako biblioteka komponentów. System wspiera pełny cykl życia zasobów (stacje, lokalizacje, EVSE) z automatycznym generowaniem punktów ładowania, zapewniając użytkownikom możliwość efektywnego organizowania infrastruktury ładowania pojazdów elektrycznych.

