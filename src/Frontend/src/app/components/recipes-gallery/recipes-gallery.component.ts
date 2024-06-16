import { Component, OnInit } from '@angular/core';
import { MealDbService } from '../../services/meal-db.service';
import { HttpResponse } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { NewRecipeComponent } from "../new-recipe/new-recipe.component";
import { Router, RouterLink, RouterLinkActive, RouterOutlet } from '@angular/router';
import { FavoritesService } from '../../services/favorites.service';
import { UsersService } from '../../services/users.service';
import { StorageService } from '../../services/storage.service';

@Component({
  selector: 'app-recipes-gallery',
  standalone: true,
  imports: [FormsModule, CommonModule, NewRecipeComponent, RouterLink],
  templateUrl: './recipes-gallery.component.html',
  styleUrl: './recipes-gallery.component.css'
})
export class RecipesGalleryComponent {
  // Ids of 9 recipes
  recipesIds = [53028, 52804, 52915, 53013, 52970, 52777, 52819, 52854, 52775];
  recipes: any = [];
  responseContent: any;
  loggedInUserId: any;

  constructor(
    private mealDbService: MealDbService,
    private favoritesService: FavoritesService,
    private storageService: StorageService,
    private userService: UsersService
  ) { }

  ngOnInit(): void {
    if (this.storageService.isLoggedIn() ) {
      var sessionUsername = this.storageService.getUsername();

      this.userService.getUsers().subscribe({
        next: (event: any) => {
          if (event instanceof HttpResponse) {
            var loggedUser = event.body.find(
              (u: { email: String, id: Number, username: String }) => 
                 u.username == sessionUsername 
            );
            this.loggedInUserId = loggedUser.id;
          }
        }
      });
    }

    // Get recipes
    this.getRecipes();
  }

  // Show 9 specific meals
  getRecipes(): any {
    for (let id of this.recipesIds) {
      const observable = this.mealDbService.getRecipeById(id);

      // Send request
      observable.subscribe({
        next: (event: any) => {
          if (event instanceof HttpResponse) {
            var recipe = event.body.meals[0];
            this.recipes.push(recipe);
          }
        },
        error: (httpErrorResponse: any) => {
          if (httpErrorResponse.error) {
            this.responseContent = httpErrorResponse.error;
          } else {
            this.responseContent = 'Une erreur est survenue !';
          }
        },
        complete: () => {
        },
      });
    }
  }

  // Add recipe to favorites
  addToFav(event: Event): any {
    // Get corresponding recipe id
    var favBtn = event.target as HTMLButtonElement;
    var recipeId = Number(favBtn.id.split('-')[1]) as number;

    if(this.storageService.isLoggedIn()) {
      var observable = this.favoritesService.addToFavorites(this.loggedInUserId, recipeId);
      observable.subscribe({
        next: (event: any) => {
          if(event instanceof HttpResponse) {
            this.responseContent = event.body.message;
            alert(`${this.responseContent}!`);
          }
        },
        error: (httpErrorResponse: any) => {
          if (httpErrorResponse.error) {
            this.responseContent = httpErrorResponse.error.message;
            alert(`${this.responseContent}!`);
          } else {
            this.responseContent = 'Une erreur est survenue !';
          }
        },
        complete: () => {
        }
      });
    }
  }
}
