import { Component, ɵSSR_CONTENT_INTEGRITY_MARKER } from '@angular/core';
import { AuthComponent } from '../auth/auth.component';
import { CommonModule } from '@angular/common';
import { RouterLink, RouterLinkActive, RouterOutlet } from '@angular/router';
import { MealDbService } from '../../services/meal-db.service';
import { FormsModule } from '@angular/forms';
import { HttpResponse } from '@angular/common/http';

@Component({
  selector: 'app-header',
  standalone: true,
  imports: [AuthComponent, CommonModule, RouterOutlet, RouterLink, RouterLinkActive, FormsModule],
  templateUrl: './header.component.html',
  styleUrl: './header.component.css'
})
export class HeaderComponent {
  searchQuery?: string;
  searchResults: string = '';

  constructor(private mealDbService: MealDbService) { }

  searchRecipe(event: Event): any {
    // console.log(this.searchQuery);
    if (this.searchQuery) {
      var observable = this.mealDbService.getRecipeByName(this.searchQuery);

      observable.subscribe({
        next: (event: any) => {
          if (event instanceof HttpResponse) {
            if (event.body.meals) {
              var meal = event.body.meals[0];
              this.searchResults = meal;
              document.location.href = `http://localhost:4200/recipes/${meal.idMeal}`
            }
            else {
              alert("Meal not found!");
            }
            console.log(this.searchResults);
          }
        },
        error: (httpErrorResponse: any) => {
          if (httpErrorResponse.error) {
            this.searchResults = httpErrorResponse.error;
          } else {
            this.searchResults = 'Une erreur est survenue !';
          }
        },
      });
    }
  }

  // Check if user is authenticated
  isUserAuthenticated(isUserAuthenticatedEvent: any): boolean {
    return isUserAuthenticatedEvent;
  }
}
