<script>
   import EnergyType from "./EnergyType.svelte";
   export let attacks;
</script>

{#if attacks}
<div>
  {#each attacks as attack}
    <div class="card-details_attack">
      <span class="title is-4">
      <span><EnergyType types={attack.cost}/></span>
        {attack.name}
      </span>
      <span>
        <span class="title is-3 has-text-muted">&nbsp;|&nbsp;</span>
        <span class="title is-4">{attack.damage}</span>
      </span>
      <p />
    </div>
  {/each}
</div>
{/if}
