<script>
   export let types;
</script>

{#if types}
{#each types as t}
<i class="energy {t.toLowerCase()}" />
{/each}
{/if}