import { Routes } from '@angular/router';
import { HomeComponent } from './components/home/home.component';
import { ContactComponent } from './components/contact/contact.component';
import { PageNotFoundComponent } from './components/page-not-found/page-not-found.component';
import { RecipesGalleryComponent } from './components/recipes-gallery/recipes-gallery.component';
import { RecipeComponent } from './components/recipe/recipe.component';
import { FavoritesComponent } from './components/favorites/favorites.component';
import { ProfileComponent } from './components/profile/profile.component';

export const routes: Routes = [
    { path: '', component: HomeComponent },
    { path: 'recipes', component: RecipesGalleryComponent },
    { path: 'favorites', component: FavoritesComponent },
    { path: 'recipes/:id', component: RecipeComponent },
    { path: 'contact', component: ContactComponent },
    { path: 'profile', component: ProfileComponent },
    { path: '**', component: PageNotFoundComponent },  // Wildcard route for a 404 page
];
