import { Injectable } from '@angular/core';

const USERNAME_KEY = 'auth-username'
const TOKEN_KEY = 'auth-token'

@Injectable({
  providedIn: 'root'
})
export class StorageService {

  constructor() { }

  cleanSession(): void {
    window.sessionStorage.clear();
  }

  public saveUser(username: string, token: string): void {
    window.sessionStorage.removeItem(USERNAME_KEY);
    window.sessionStorage.removeItem(TOKEN_KEY);
    
    window.sessionStorage.setItem(USERNAME_KEY, JSON.stringify(username));
    window.sessionStorage.setItem(TOKEN_KEY, JSON.stringify(token));
  }

  public getUsername(): any {
    const user = window.sessionStorage.getItem(USERNAME_KEY);
    if (user) {
      return JSON.parse(user);
    }

    return {};
  }

  public getUserToken(): any {
    const token = window.sessionStorage.getItem(TOKEN_KEY);
    if(token) {
      return JSON.parse(token);
    }

    return {};
  }

  public isLoggedIn(): boolean {
    const userToken = window.sessionStorage.getItem(TOKEN_KEY);
    if (userToken) {
      return true;
    }

    return false;
  }
}
