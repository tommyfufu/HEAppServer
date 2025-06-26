import { Component, Input, OnChanges, SimpleChanges } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ApiService } from '../api-service.service';
import { Record } from '../models/patient.model';
import { NgIf, NgFor } from '@angular/common';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-game-record',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './game-record.component.html',
  styleUrl: './game-record.component.css'
})

export class GameRecordComponent {
  @Input() userId?: string;
  records: Record[] = [];
  filteredRecords: Record[] = [];

  gameIds: number[] = [];
  selectedGameId: string = '';
  gameSelected: boolean = false;

  loading = false;
  error: string | null = null;

  gameList = [
    { id: '0', name: '數字點點名' },
    { id: '1', name: '五顏配六色' },
    { id: '2', name: '按鈕排排站' },
  ];

  constructor(private api: ApiService) {}

  ngOnChanges(changes: SimpleChanges) {
    if (changes['userId'] && this.userId) {
      this.fetchRecords(this.userId);
    }
  }

  private fetchRecords(userId: string) {
    this.loading = true;
    this.error = null;

    this.api.getRecords(userId).subscribe({
      next: (recs) => {
        this.records = recs.reverse();
        this.filteredRecords = [];
        this.gameIds = Array.from(new Set(recs.map(r => r.gameId))); // 去重後的 Game ID
        this.loading = false;
      },
      error: (err) => {
        console.error('Failed to load records', err);
        this.error = '無法讀取遊戲紀錄';
        this.loading = false;
      }
    });
  }

  onSelectGame(gameId: string) {
    this.selectedGameId = gameId;
    this.gameSelected = true;
    if (gameId === '') {
      this.filteredRecords = [];
    } else {
      this.filteredRecords = this.records.filter(r => r.gameId.toString() === gameId);
    }
  }
}
