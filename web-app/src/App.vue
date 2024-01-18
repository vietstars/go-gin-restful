<template>
  <the-header></the-header>
  <the-filters @fetch-movies="handleFilters" ></the-filters>
  <PagesControll @page-control="fetchByPage" ></PagesControll>
  <LoadingSpinner v-if="loading" />
  <h1 v-else-if="!loading && errors" class="text-center font-bold text-4xl my-12">Something went wrong - try again later!</h1>
  <movies-list v-else :movies="moviesData" ></movies-list>
</template>

<script>


import TheHeader from './components/layout/TheHeader.vue'
import TheFilters from './components/layout/TheFilters.vue'
import MoviesList from './components/MoviesList.vue';
import LoadingSpinner from './components/layout/LoadingSpinner.vue';
import PagesControll from './components/layout/PagesControll.vue';
import { API_KEY } from './config/constants';

export default {
  name: 'App',
  components: {
    TheHeader,
    TheFilters,
    MoviesList,
    LoadingSpinner,
    PagesControll
},
  data() {
    return {
      moviesData: [],
      errors: null,
      loading: false,
      offsetNum: 0,
      movieOrder: "by-opening-date",
      query: ""
    }
  },
  methods: {
    handleFilters(movieOrder, search) {
      this.fetchMovies(this.offsetNum, movieOrder, search);
    },
    fetchByPage(offsetNum) {
      this.offsetNum = offsetNum;
      this.fetchMovies(this.offsetNum, this.movieOrder, this.query);
    },
    async fetchMovies(offsetNum = 0, movieOrder = "by-opening-date", query = "") {
        this.errors = null;
        this.loading = true;
        try {
          this.moviesData = [
            {
              display_title: 'title 123',
              link:{
                url:'url-123',
              },
              summary_short: 'summary_short 123',
              multimedia: {
                src: 'https://picsum.photos/400/200',
              },
              publication_date: '2024-12-10',
            },
            {
              display_title: 'title 456',
              link:{
                url:'url-456',
              },
              summary_short: 'summary_short 456',
              multimedia: {
                src: 'https://picsum.photos/400/200',
              },
              publication_date: '2024-12-10',
            },
            {
              display_title: 'title  789',
              link:{
                url:'url-789',
              },
              summary_short: 'summary_short 789',
              multimedia: {
                src: 'https://picsum.photos/400/200',
              },
              publication_date: '2024-12-10',
            },
            {
              display_title: 'title  567',
              link:{
                url:'url-567',
              },
              summary_short: 'summary_short 567',
              multimedia: {
                src: 'https://picsum.photos/400/200',
              },
              publication_date: '2024-12-10',
            },
            {
              display_title: 'title  234',
              link:{
                url:'url-234',
              },
              summary_short: 'summary_short 234',
              multimedia: {
                src: 'https://picsum.photos/400/200',
              },
              publication_date: '2024-12-10',
            },
            {
              display_title: 'title  345',
              link:{
                url:'url-345',
              },
              summary_short: 'summary_short 345',
              multimedia: {
                src: 'https://picsum.photos/400/200',
              },
              publication_date: '2024-12-10',
            },
          ];

          this.loading = false;
            // const response = await fetch(`https://api.nytimes.com/svc/movies/v2/reviews/picks.json?api-key=${API_KEY}&offset=${offsetNum}&order=${movieOrder}&query=${query}`);
            // if (!response.ok) {
            //     const error = new Error("Failed to fetch Data");
            //     error.statusCode = 404;
            //     throw error;
            // }
            // else {
            //     const data = await response.json();
            //     this.moviesData = data.results;
            //     this.loading = false;
            // }
        } catch (error) {
            console.log(error);
            this.loading = false;
            this.errors = "Something went wrong - try again later!";
        }
    }
  },
  async mounted() {
        document.title = "NYT Critic's Picks";
       await this.fetchMovies();
  }
}
</script>