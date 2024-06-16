import { HttpClient, HttpHeaders, HttpRequest } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BaseRouteReuseStrategy } from '@angular/router';
import { Observable } from 'rxjs';

// const BASE_URL = "http://localhost:8585/favoriteRecipes"
const BASE_URL = "https://14oi3yhq0g.execute-api.ca-central-1.amazonaws.com/favoriteRecipes"
const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' })
};

@Injectable({
  providedIn: 'root'
})
export class FavoritesService {

  constructor(private http: HttpClient) { }

  // Add a recipe to the list of user favorite recipes
  addToFavorites(userId: number, recipeId: number): Observable<any> {
    const body = {
      "userId": userId,
      "recipeId": recipeId
    };

    const postReq = new HttpRequest("POST", BASE_URL, body, httpOptions);

    return this.http.request(postReq);
  }

  // Get all favorite recipes for a specifi user (id) 
  getFavoriteRecipes(userId: number): Observable<any> {
    const getReq = new HttpRequest("GET", `${BASE_URL}/${userId}`);

    return this.http.request(getReq);
  }

  // Remove a favorite from the list of the user favorite recipes
  removeFavorite(recipeId: number): Observable<any> {
    const getReq = new HttpRequest("DELETE", `${BASE_URL}/${recipeId}`);

    return this.http.request(getReq);
  }
}
