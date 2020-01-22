<script>
  import { onMount } from "svelte";
  import Attack from "./Attack.svelte";
  import EnergyType from "./EnergyType.svelte";

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

<h1>Pokedex</h1>

<div class="container">
  {#each cards as card}
    <div class="columns">
      <div class="column is-one-third">
        <img class="card-img" src={card.imageUrl} alt={card.name} />
      </div>
      <div class="column is-6 is-offset-1">
        <div class="content card-details">
          <nav class="card-details_head level">
            <div class="level-left">
              <div class="level-item">
                <span class="title is-3">{card.name}</span>
              </div>
              <div class="level-item">
                <span class="title is-4 has-text-muted">
                  {card.supertype} - {card.subtype}
                </span>
              </div>
            </div>
            <div class="level-right">
              <div class="level-item">
                <span class="title is-5">
                  HP {card.hp}
                  <span>
                    <EnergyType types={card.types}/>
                  </span>
                </span>
              </div>
            </div>
          </nav>
          <hr />
          <Attack attacks={card.attacks} />
          <nav class="level card-details_weakness-resistance-retreat">
            <div class="level-item has-text-centered">
              <div class="card-details_weakness">
                <p class="heading">weakness</p>
                <p class="title is-5">
                  <span>
                    <span>
                      <i class="energy water" />
                    </span>
                    +10
                  </span>
                </p>
              </div>
            </div>
            <div class="level-item has-text-centered">
              <div class="card-details_weakness">
                <p class="heading">resistances</p>
                <p class="title is-5">
                  <span>
                    <span>
                      <i class="energy lightning" />
                    </span>
                    -20
                  </span>
                </p>
              </div>
            </div>
            <div class="level-item has-text-centered">
              <div class="card-details_weakness">
                <p class="heading">retreat cost</p>
                <p class="title is-5">
                  <span>
                    <i class="energy colorless" />
                  </span>
                </p>
              </div>
            </div>
          </nav>
          <nav class="level">
            <div class="level-item has-text-centered">
              <div class="card-details_artist">
                <p class="heading">artist</p>
                <p class="title is-5">Kagemaru Himeno</p>
              </div>
            </div>
            <div class="level-item has-text-centered">
              <div class="card-details_artist">
                <p class="heading">rarity</p>
                <p class="title is-5">Common</p>
              </div>
            </div>
            <div class="level-item has-text-centered">
              <div class="card-details_artist">
                <p class="heading">set</p>
                <p class="title is-5">
                  Legends Awakened
                  <img
                    src="https://images.pokemontcg.io/dp6/symbol.png"
                    alt="set"
                    width="24"
                    height="24" />
                </p>
              </div>
            </div>
          </nav>
        </div>
      </div>
    </div>
  {:else}
    <!-- this block renders when cards.length === 0 -->
    <p>loading...</p>
  {/each}
</div>
