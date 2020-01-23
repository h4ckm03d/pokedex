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
      `https://api.pokemontcg.io/v1/cards?pageSize=10&name=${query}`
    );
    const payload = await res.json();
    cards = payload.cards;
  }

  function searchClick() {
    console.log("klicked", query)
    fetchData();
  }
</script>

<style>
  .container {
    display: grid;
    /* grid-template-columns: repeat(auto-fit, minmax(245px, 1fr));
   grid-gap: 20px; */
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
  <div class="field has-addons">
    <p class="control">
      <input
        id="search-input-text"
        class="input"
        bind:value={query}
        placeholder="Filter by name, types" />
    </p>
    <p class="control">
    <button class="button" on:click={searchClick}>
      <span class="icon icon-input-search">
        <i class="fas fa-search" />
      </span>
      <span>Search</span>
    </button>
  </p>
  </div>
  {#each cards as card}
    <CardInfo card={card} />
    <hr />
  {:else}
    <!-- this block renders when cards.length === 0 -->
    <p>loading...</p>
  {/each}
</div>
