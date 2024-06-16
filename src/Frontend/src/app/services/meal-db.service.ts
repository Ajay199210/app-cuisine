import { HttpClient, HttpRequest } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class MealDbService {
  private baseUrl = "https://www.themealdb.com/api/json/v1/1"

  constructor(private http: HttpClient) { }

  // Get a random recipes
  getRandomMeal(): Observable<any> {
    const getReq = new HttpRequest("GET", `${this.baseUrl}/random.php`)

    return this.http.request(getReq);
  }

  // Get a recipe by id
  getRecipeById(id: number): Observable<any> {
    const getReq = new HttpRequest("GET", `${this.baseUrl}/lookup.php?i=${id}`)

    return this.http.request(getReq);
  }

  // Get a recipe by name
  getRecipeByName(recipeName: string): Observable<any> {
    const getReq = new HttpRequest("GET", `${this.baseUrl}/search.php?s=${recipeName}`)

    return this.http.request(getReq);
  }
}
