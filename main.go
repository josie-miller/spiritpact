package main

import (
    "html/template"
    "log"
    "net/http"
    "time"

    "github.com/gorilla/mux"
    "github.com/gorilla/sessions"
)

var (
    // Templates cache
    templates *template.Template
    store     = sessions.NewCookieStore([]byte("super-secret-key"))
)

func init() {
    // Load all templates
    templates = template.Must(template.ParseGlob("templates/*.html"))
}

// Utility to render templates
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    err := templates.ExecuteTemplate(w, tmpl, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// Home Page
func homeHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "home.html", nil)
}

// About Us Page
func aboutHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "about.html", nil)
}

// Blog Page
func blogHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "blog.html", nil)
}

// Partners Page
func partnersHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "partners.html", nil)
}

// Bookings Page
func bookingsHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "bookings.html", nil)
}

// Event Details Page
func eventDetailsHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "event_details.html", nil)
}

// Login Page
func loginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        username := r.FormValue("username")
        password := r.FormValue("password")

        // Simulate authentication
        if username == "user" && password == "password" {
            session, _ := store.Get(r, "session")
            session.Values["authenticated"] = true
            session.Save(r, w)
            http.Redirect(w, r, "/dashboard", http.StatusFound)
        } else {
            http.Error(w, "Invalid login credentials", http.StatusUnauthorized)
        }
    }
    renderTemplate(w, "login.html", nil)
}

// Registration Page
func registerHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        username := r.FormValue("username")
        password := r.FormValue("password")

        // Simulate user registration
        log.Printf("Registered User: %s\n", username)
        http.Redirect(w, r, "/login", http.StatusFound)
    }
    renderTemplate(w, "register.html", nil)
}

// User Dashboard
func dashboardHandler(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "session")

    // Check if user is authenticated
    if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    renderTemplate(w, "user_dashboard.html", nil)
}

// Profile Page
func profileHandler(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "session")

    if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    renderTemplate(w, "profile.html", nil)
}

// List Bookings (for authenticated users)
func bookingListHandler(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "session")

    if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    renderTemplate(w, "booking_list.html", nil)
}

// Manage Bookings
func manageBookingHandler(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "session")

    if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    renderTemplate(w, "manage_booking.html", nil)
}

func main() {
    r := mux.NewRouter()

    // Serve static assets
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

    // Routes
    r.HandleFunc("/", homeHandler)
    r.HandleFunc("/about", aboutHandler)
    r.HandleFunc("/blog", blogHandler)
    r.HandleFunc("/partners", partnersHandler)
    r.HandleFunc("/bookings", bookingsHandler)
    r.HandleFunc("/event/{id}", eventDetailsHandler)
    r.HandleFunc("/login", loginHandler)
    r.HandleFunc("/register", registerHandler)
    r.HandleFunc("/dashboard", dashboardHandler)
    r.HandleFunc("/profile", profileHandler)
    r.HandleFunc("/my-bookings", bookingListHandler)
    r.HandleFunc("/manage-bookings", manageBookingHandler)

    // Start the server
    srv := &http.Server{
        Handler:      r,
        Addr:         ":8080",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }
    log.Println("Starting server on :8080")
    log.Fatal(srv.ListenAndServe())
}
