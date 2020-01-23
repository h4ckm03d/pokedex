<script>
  import { onMount } from "svelte";
  import CardInfo from "./CardInfo.svelte";
  let cards = [];

  onMount(async () => {
    const res = await fetch(`https://api.pokemontcg.io/v1/cards?pageSize=10`);
    console.log(res);
    const payload = await res.json();
    cards = payload.cards;
  });
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
      <h1 class="title">
        Pokedex
      </h1>
      <h2 class="subtitle">
        Pokemon TCG
      </h2>
    </div>
  </div>
</section>
<hr />
<div class="container">
  {#each cards as card}
    <CardInfo {card} />
    <hr />
  {:else}
    <!-- this block renders when cards.length === 0 -->
    <p>loading...</p>
  {/each}
</div>
