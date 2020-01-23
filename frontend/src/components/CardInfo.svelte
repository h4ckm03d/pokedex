<script>
  import Attack from "./Attack.svelte";
  import Ability from "./Ability.svelte";
  import EnergyType from "./EnergyType.svelte";
  import Effectiveness from "./Effectiveness.svelte";
  export let card;
</script>

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
                <EnergyType types={card.types} />
              </span>
            </span>
          </div>
        </div>
      </nav>
      <hr />
      <Ability ability={card.ability} />
      <Attack attacks={card.attacks} />
      <nav class="level card-details_weakness-resistance-retreat">
        <div class="level-item has-text-centered">
          <div class="card-details_weakness">
            <p class="heading">weakness</p>
            <Effectiveness effects={card.weaknesses} />
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div class="card-details_weakness">
            <p class="heading">resistances</p>
            <Effectiveness effects={card.resistances} />
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div class="card-details_weakness">
            <p class="heading">retreat cost</p>
            <p class="title is-5">
              <span>
                <EnergyType types={card.retreatCost} />
              </span>
            </p>
          </div>
        </div>
      </nav>
      <nav class="level">
        <div class="level-item has-text-centered">
          <div class="card-details_artist">
            <p class="heading">artist</p>
            <p class="title is-5">{card.artist}</p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div class="card-details_artist">
            <p class="heading">rarity</p>
            <p class="title is-5">{card.rarity}</p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div class="card-details_artist">
            <p class="heading">set</p>
            <p class="title is-5">
              {card.set}
              <img
                src="https://images.pokemontcg.io/{card.setCode}/symbol.png"
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
