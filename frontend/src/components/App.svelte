<script>
  import { onMount } from "svelte";
  import CardInfo from "./CardInfo.svelte";
  let cards = [];
  let query = "";
  let page;

  onMount(async () => {
    fetchData();
  });

  async function fetchData() {
    const res = await fetch(
      `https://0e8hk047l4.execute-api.ap-southeast-1.amazonaws.com/demo?pageSize=10&name=${query}`,
      {
        method: "GET", // *GET, POST, PUT, DELETE, etc.
        mode: 'cors'
      }
    );
    cards = await res.json();
  }
  function enterSearch(event) {
    // Number 13 is the "Enter" key on the keyboard
    if (event.keyCode === 13) {
      // Cancel the default action, if needed
      event.preventDefault();
      fetchData();
    }
  }

  function searchClick() {
    fetchData();
  }
</script>

<style>
  .container {
    display: grid;
  }
</style>

<section class="hero">
  <div class="hero-body">
    <div class="container">
      <h1 class="title">Pokedex</h1>
      <h2 class="subtitle">Pokemon TCG</h2>
    </div>
  </div>
</section>
<hr />
<div class="container">
  <div class="column">
    <div class="field has-addons">
      <p class="control is-expanded">
        <input
          id="search-input-text"
          class="input"
          type="text"
          bind:value={query}
          on:keyup={enterSearch}
          placeholder="Filter by name" />
      </p>
      <p class="control">
        <button class="button is-primary" on:click={searchClick}>
          <span>Search</span>
        </button>
      </p>
    </div>
  </div>
  {#each cards as card}
    <CardInfo {card} />
    <hr />
  {:else}
    <!-- this block renders when cards.length === 0 -->
    <p>loading...</p>
  {/each}
</div>
