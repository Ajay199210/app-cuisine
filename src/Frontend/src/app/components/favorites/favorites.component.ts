import { Component, OnInit } from '@angular/core';
import { FavoritesService } from '../../services/favorites.service';
import { StorageService } from '../../services/storage.service';
import { UsersService } from '../../services/users.service';
import { HttpResponse } from '@angular/common/http';
import { MealDbService } from '../../services/meal-db.service';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'app-favorites',
  standalone: true,
  imports: [RouterLink],
  templateUrl: './favorites.component.html',
  styleUrl: './favorites.component.css'
})
export class FavoritesComponent implements OnInit {
  loggedInUserId: any;
  userFavorites: any[] = []; // recipes objects
  userFavoritesCount = 0;
  responseContent: any;
  favIds: any[] = []; // favorites Ids

  constructor(private favoritesService: FavoritesService, private userService: UsersService,
    private storageService: StorageService, private mealDbService: MealDbService) { }

  ngOnInit(): void {
    if (this.storageService.isLoggedIn()) {
      var sessionUsername = this.storageService.getUsername();

      this.userService.getUsers().subscribe({
        next: (event: any) => {
          if (event instanceof HttpResponse) {
            var loggedUser = event.body.find(
              (u: { email: String, id: Number, username: String }) =>
                u.username == sessionUsername
            );
            this.loggedInUserId = loggedUser.id;
            this.getUserFavoriteRecipes();
          }
        }
      });
    }
  }

  // Show user favorite recipes
  getUserFavoriteRecipes(): void {
    const favoritesObservable = this.favoritesService.getFavoriteRecipes(this.loggedInUserId);

    // Send request
    favoritesObservable.subscribe({
      next: (event: any) => {
        if (event instanceof HttpResponse) {
          this.favIds = event.body.map(
            (fav : { ID: Number, userId: Number, recipeId: Number }) => fav.ID
          );
          for (const fav of event.body) {
            var favId = fav.recipeId;
            var mealDbObservable = this.mealDbService.getRecipeById(favId);

            mealDbObservable.subscribe({
              next: (event: any) => {
                if (event instanceof HttpResponse) {
                  var userFavRecipe = event.body.meals[0];
                  this.userFavorites.push(userFavRecipe);
                  this.userFavoritesCount++;
                }
              }
            });
          }
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

  // Remove recipe from user favorites
  removeRecipe(event: any): void {
    var favRecipeId = event.target.id.split('-')[1];

    var observable = this.favoritesService.removeFavorite(favRecipeId);
    
    observable.subscribe({
      next: (event: any) => {
        if(event instanceof HttpResponse) {
          if(event.ok) {
            window.location.reload();
          }
        }
      },
      error: (httpErrorResponse: any) => {
        if (httpErrorResponse.error) {
          this.responseContent = httpErrorResponse.error.message;
        } else {
          this.responseContent = 'Une erreur est survenue !';
        }
      },
      complete: () => {
      },
    });
  }
}
