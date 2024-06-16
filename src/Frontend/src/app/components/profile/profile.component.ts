import { Component, OnInit } from '@angular/core';
import { StorageService } from '../../services/storage.service';
import { UsersService } from '../../services/users.service';
import { HttpResponse } from '@angular/common/http';

@Component({
  selector: 'app-profile',
  standalone: true,
  imports: [],
  templateUrl: './profile.component.html',
  styleUrl: './profile.component.css'
})
export class ProfileComponent implements OnInit {
  isUserAuthenticated: boolean = false;
  username: string = '';

  constructor(private storageService: StorageService,
    private userService: UsersService) { }

  ngOnInit(): void {
    this.isUserAuthenticated = this.storageService.isLoggedIn();

    if (this.isUserAuthenticated) {
      this.username = this.storageService.getUsername();
    }
  }
}
