// import { Component, OnInit } from '@angular/core';
// import { ApiService } from '../api-service.service';

// @Component({
//   selector: 'app-bulletin-board',
//   templateUrl: './bulletin-board.component.html',
//   standalone: true,
//   styleUrls: ['./bulletin-board.component.css']
// })
// export class BulletinBoardComponent implements OnInit {
//   messages: Record<string, string>[] = []; 

//   constructor(private apiService: ApiService) { }

//   ngOnInit(): void {
//     this.apiService.getMessagesForPatient(patientID).subscribe({
//       next: (msgs: Record<string, string>) => {
//         this.messages = msgs;
//       },
//       error: (error) => {
//         console.error('Error fetching messages:', error);
//       }
//     });
//   }
// }