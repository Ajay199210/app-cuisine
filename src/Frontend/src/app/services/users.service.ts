import { HttpClient, HttpHeaders, HttpRequest } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class UsersService {
  // private baseUrl = "http://localhost:9090"
  private baseUrl = "https://14oi3yhq0g.execute-api.ca-central-1.amazonaws.com"

  constructor(private http: HttpClient) { }

  // Afficher les utilisateurs
  getUsers(): Observable<any> {
    const getReq = new HttpRequest("GET", `${this.baseUrl}/users`)

    return this.http.request(getReq);
  }

  // Authenticate user
  loginUser(username: string, password: string): Observable<any> {
    const headers = new HttpHeaders();
    headers.set('Content-Type', 'application/json');

    const body = {
      "username": username,
      "password": password
    };

    const postReq = new HttpRequest("POST", `${this.baseUrl}/login`, body, { headers });

    return this.http.request(postReq);
  }

  // Logout
  logoutUser(): Observable<any> {
    return this.http.get(`${this.baseUrl}/logout`);
  }

  // Créer un nouveau compte
  registerUser(username: string, email: string,
    password: string, confirmPassword: string): Observable<any> {

    const headers = new HttpHeaders();
    headers.set('Content-Type', 'application/json');

    const body = {
      "email": email,
      "username": username,
      "password": password,
      "passwordconfirmed": confirmPassword
    }

    const postReq = new HttpRequest("POST", `${this.baseUrl}/signup`, body, { headers });

    return this.http.request(postReq);
  }
}
