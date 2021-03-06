# Routing data for GraphHopper

Routing data for [GraphHopper](https://www.graphhopper.com) for offline navigation. With one-click installation on the best outdoor navigation app [Locus Map](https://www.locusmap.eu).

File [issue](https://github.com/develar/gh-routing-data/issues) if routing data not provided for wanted country or region.

!!! tip "Cross border navigation"
    As GraphHopper doesn't support [multiple](https://github.com/graphhopper/graphhopper/issues/293) files, if you need to cross borders, please use special region wide routing data (e.g. Alps). Feel free to file [issue](https://github.com/develar/gh-routing-data/issues) to build a special region for your needs.

## Installation

Click a "Locus" link to install on Locus (will be automatically downloaded and extracted to `mapsVector/`).
In the [GraphHopper Add-on](https://github.com/asamm/locus-addon-graphhopper/releases/latest) choose which file you want to use.

Or simply download zip file to install manually.
As most zip libraries for Android doesn't support files more than 2GB, large regions split into 3 parts.

Click on the region name to see other download options.

## Maps

!!! question "Which vehicles are supported?"
    * pedestrian or walking with priority for more beautiful hiking tours (`hike`),
    * trekking bike avoiding hills (`bike2`),
    * mountain bike,
    * racing bike,
    * car.
    
    Foot routing not supported to reduce size of graph (please use `hike` instead).
    
??? question "How often data is updated?"
    At least monthly. 
    File [issue](https://github.com/develar/gh-routing-data/issues) to force update if need. 
    Also, data can be updated once a new version of GraphHopper is released.
    
    Last update: 2020-03-02.

<!-- do not edit. start of generated block -->

<label for="ghVersions">Locus Map Add-on Version:</label>
<select name="ghVersions" id="mapVersionFormatSelect">
  <option value="1.0-pre20">0.9 | GraphHopper 1.0-pre20</option>
  <option value="1.0-pre26">0.10 (unreleased) | GraphHopper 1.0-pre26</option>
  <option value="1.0-pre31">0.10 (unreleased) | GraphHopper 1.0-pre31</option>
</select>
<div>
  <small>[Europe](#europe) | [Northern Europe](#northern-europe) | [North America](#north-america) | [Asia](#asia) | [Other](#other)</small>
</div>

### Europe

<div class="v-1.0-pre20">
  
<table>
<tr>
  <th>Region</th>
  <th>Install</th>
  <th>Size</th>
  <th>Coverage</th>
</tr>

<tr>
  <td class="regionInfo">Alps</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/alps.locus.xml">Locus</a></td>
  <td>1.7 GB</td>
  <td><a href="/coverage.html#alps@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/alps.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Bayern (Germany), Austria, Czech Republic</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/bayern-at-cz.locus.xml">Locus</a></td>
  <td>1.2 GB</td>
  <td><a href="/coverage.html#bayern-at-cz@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/bayern-at-cz.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Belgium</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/belgium.locus.xml">Locus</a></td>
  <td>216 MB</td>
  <td><a href="/coverage.html#belgium@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/belgium.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Czech Republic</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/czech-republic.locus.xml">Locus</a></td>
  <td>273 MB</td>
  <td><a href="/coverage.html#czech-republic@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/czech-republic.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Estonia, Latvia and Lithuania</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/estonia-latvia-lithuania.locus.xml">Locus</a></td>
  <td>174 MB</td>
  <td><a href="/coverage.html#estonia-latvia-lithuania@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/estonia-latvia-lithuania.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">France</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/france.locus.xml">Locus</a></td>
  <td>2.0 GB</td>
  <td><a href="/coverage.html#france@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/france.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Germany</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/germany.locus.xml">Locus</a></td>
  <td>2.5 GB</td>
  <td><a href="/coverage.html#germany@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip:
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/germany-part1.osm-gh.zip">part 1</a> (669 MB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/germany-part2.osm-gh.zip">part 2</a> (594 MB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/germany-part3.osm-gh.zip">part 3</a> (1.3 GB)
    
</td>
</tr>

<tr>
  <td class="regionInfo">Germany, Austria and Switzerland</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/de-at-ch.locus.xml">Locus</a></td>
  <td>3.2 GB</td>
  <td><a href="/coverage.html#dach@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip:
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/de-at-ch-part1.osm-gh.zip">part 1</a> (820 MB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/de-at-ch-part2.osm-gh.zip">part 2</a> (734 MB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/de-at-ch-part3.osm-gh.zip">part 3</a> (1.6 GB)
    
</td>
</tr>

<tr>
  <td class="regionInfo">Greece</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/greece.locus.xml">Locus</a></td>
  <td>281 MB</td>
  <td><a href="/coverage.html#greece@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/greece.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Italy</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/italy.locus.xml">Locus</a></td>
  <td>1.2 GB</td>
  <td><a href="/coverage.html#italy@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/italy.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Netherlands</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/netherlands.locus.xml">Locus</a></td>
  <td>372 MB</td>
  <td><a href="/coverage.html#netherlands@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/netherlands.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Poland</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/poland.locus.xml">Locus</a></td>
  <td>818 MB</td>
  <td><a href="/coverage.html#poland@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/poland.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Portugal and Spain</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/portugal-spain.locus.xml">Locus</a></td>
  <td>1.4 GB</td>
  <td><a href="/coverage.html#portugal-spain@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/portugal-spain.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Russia</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/russia.locus.xml">Locus</a></td>
  <td>1.8 GB</td>
  <td><a href="/coverage.html#russia@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/russia.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Switzerland</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/switzerland.locus.xml">Locus</a></td>
  <td>266 MB</td>
  <td><a href="/coverage.html#switzerland@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/switzerland.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Ukraine</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/ukraine.locus.xml">Locus</a></td>
  <td>474 MB</td>
  <td><a href="/coverage.html#ukraine@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/ukraine.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Albania, Bosnia-Herzegovina, Bulgaria, Croatia, Hungary, Kosovo, Macedonia, Moldova, Montenegro, Romania, Serbia, Slovakia and Slovenia</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/al-ba-bg-hr-hu-xk-mk-md-me-ro-rs-sk-si.locus.xml">Locus</a></td>
  <td>1.0 GB</td>
  <td><a href="/coverage.html#al-ba-bg-hr-hu-xk-mk-md-me-ro-rs-sk-si@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/al-ba-bg-hr-hu-xk-mk-md-me-ro-rs-sk-si.osm-gh.zip">download</a>
</td>
</tr>
 
</table>
</div>

<div class="v-1.0-pre26" style="display: none">
  
<table>
<tr>
  <th>Region</th>
  <th>Install</th>
  <th>Size</th>
  <th>Coverage</th>
</tr>

<tr>
  <td class="regionInfo">Alps</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/alps.locus.xml">Locus</a></td>
  <td>1.7 GB</td>
  <td><a href="/coverage.html#alps@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/alps.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Austria</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/austria.locus.xml">Locus</a></td>
  <td>406 MB</td>
  <td><a href="/coverage.html#austria@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/austria.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Bayern (Germany), Austria, Czech Republic</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/bayern-at-cz.locus.xml">Locus</a></td>
  <td>1.2 GB</td>
  <td><a href="/coverage.html#bayern-at-cz@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/bayern-at-cz.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Belgium</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/belgium.locus.xml">Locus</a></td>
  <td>219 MB</td>
  <td><a href="/coverage.html#belgium@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/belgium.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Czech Republic</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/czech-republic.locus.xml">Locus</a></td>
  <td>275 MB</td>
  <td><a href="/coverage.html#czech-republic@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/czech-republic.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Estonia, Latvia and Lithuania</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/estonia-latvia-lithuania.locus.xml">Locus</a></td>
  <td>176 MB</td>
  <td><a href="/coverage.html#estonia-latvia-lithuania@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/estonia-latvia-lithuania.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">France</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/france.locus.xml">Locus</a></td>
  <td>2.0 GB</td>
  <td><a href="/coverage.html#france@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/france.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Germany</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/germany.locus.xml">Locus</a></td>
  <td>2.6 GB</td>
  <td><a href="/coverage.html#germany@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip:
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/germany-part1.osm-gh.zip">part 1</a> (680 MB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/germany-part2.osm-gh.zip">part 2</a> (602 MB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/germany-part3.osm-gh.zip">part 3</a> (1.3 GB)
    
</td>
</tr>

<tr>
  <td class="regionInfo">Germany, Austria and Switzerland</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/de-at-ch.locus.xml">Locus</a></td>
  <td>3.2 GB</td>
  <td><a href="/coverage.html#dach@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip:
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/de-at-ch-part1.osm-gh.zip">part 1</a> (832 MB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/de-at-ch-part2.osm-gh.zip">part 2</a> (743 MB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/de-at-ch-part3.osm-gh.zip">part 3</a> (1.7 GB)
    
</td>
</tr>

<tr>
  <td class="regionInfo">Greece</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/greece.locus.xml">Locus</a></td>
  <td>284 MB</td>
  <td><a href="/coverage.html#greece@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/greece.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Italy</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/italy.locus.xml">Locus</a></td>
  <td>1.2 GB</td>
  <td><a href="/coverage.html#italy@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/italy.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Netherlands</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/netherlands.locus.xml">Locus</a></td>
  <td>375 MB</td>
  <td><a href="/coverage.html#netherlands@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/netherlands.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Poland</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/poland.locus.xml">Locus</a></td>
  <td>825 MB</td>
  <td><a href="/coverage.html#poland@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/poland.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Portugal and Spain</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/portugal-spain.locus.xml">Locus</a></td>
  <td>1.4 GB</td>
  <td><a href="/coverage.html#portugal-spain@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/portugal-spain.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Russia</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/russia.locus.xml">Locus</a></td>
  <td>1.8 GB</td>
  <td><a href="/coverage.html#russia@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/russia.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Switzerland</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/switzerland.locus.xml">Locus</a></td>
  <td>270 MB</td>
  <td><a href="/coverage.html#switzerland@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/switzerland.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Ukraine</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/ukraine.locus.xml">Locus</a></td>
  <td>479 MB</td>
  <td><a href="/coverage.html#ukraine@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/ukraine.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Albania, Bosnia-Herzegovina, Bulgaria, Croatia, Hungary, Kosovo, Macedonia, Moldova, Montenegro, Romania, Serbia, Slovakia and Slovenia</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/al-ba-bg-hr-hu-xk-mk-md-me-ro-rs-sk-si.locus.xml">Locus</a></td>
  <td>1.0 GB</td>
  <td><a href="/coverage.html#al-ba-bg-hr-hu-xk-mk-md-me-ro-rs-sk-si@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/al-ba-bg-hr-hu-xk-mk-md-me-ro-rs-sk-si.osm-gh.zip">download</a>
</td>
</tr>
 
</table>
</div>

<div class="v-1.0-pre31" style="display: none">
  
<table>
<tr>
  <th>Region</th>
  <th>Install</th>
  <th>Size</th>
  <th>Coverage</th>
</tr>

<tr>
  <td class="regionInfo">Alps</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/alps.locus.xml">Locus</a></td>
  <td>1.7 GB</td>
  <td><a href="/coverage.html#alps@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/alps.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Austria</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/austria.locus.xml">Locus</a></td>
  <td>404 MB</td>
  <td><a href="/coverage.html#austria@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/austria.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Bayern (Germany), Austria, Czech Republic</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/bayern-at-cz.locus.xml">Locus</a></td>
  <td>1.2 GB</td>
  <td><a href="/coverage.html#bayern-at-cz@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/bayern-at-cz.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Belgium</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/belgium.locus.xml">Locus</a></td>
  <td>216 MB</td>
  <td><a href="/coverage.html#belgium@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/belgium.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Czech Republic</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/czech-republic.locus.xml">Locus</a></td>
  <td>273 MB</td>
  <td><a href="/coverage.html#czech-republic@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/czech-republic.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Estonia, Latvia and Lithuania</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/estonia-latvia-lithuania.locus.xml">Locus</a></td>
  <td>175 MB</td>
  <td><a href="/coverage.html#estonia-latvia-lithuania@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/estonia-latvia-lithuania.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">France</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/france.locus.xml">Locus</a></td>
  <td>2.0 GB</td>
  <td><a href="/coverage.html#france@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/france.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Germany</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/germany.locus.xml">Locus</a></td>
  <td>2.5 GB</td>
  <td><a href="/coverage.html#germany@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip:
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/germany-part1.osm-gh.zip">part 1</a> (1.2 GB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/germany-part2.osm-gh.zip">part 2</a> (604 MB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/germany-part3.osm-gh.zip">part 3</a> (776 MB)
    
</td>
</tr>

<tr>
  <td class="regionInfo">Germany, Austria and Switzerland</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/de-at-ch.locus.xml">Locus</a></td>
  <td>3.2 GB</td>
  <td><a href="/coverage.html#dach@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip:
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/de-at-ch-part1.osm-gh.zip">part 1</a> (1.4 GB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/de-at-ch-part2.osm-gh.zip">part 2</a> (747 MB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/de-at-ch-part3.osm-gh.zip">part 3</a> (1.0 GB)
    
</td>
</tr>

<tr>
  <td class="regionInfo">Greece</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/greece.locus.xml">Locus</a></td>
  <td>286 MB</td>
  <td><a href="/coverage.html#greece@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/greece.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Italy</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/italy.locus.xml">Locus</a></td>
  <td>1.2 GB</td>
  <td><a href="/coverage.html#italy@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/italy.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Netherlands</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/netherlands.locus.xml">Locus</a></td>
  <td>374 MB</td>
  <td><a href="/coverage.html#netherlands@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/netherlands.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Poland</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/poland.locus.xml">Locus</a></td>
  <td>817 MB</td>
  <td><a href="/coverage.html#poland@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/poland.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Portugal and Spain</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/portugal-spain.locus.xml">Locus</a></td>
  <td>1.4 GB</td>
  <td><a href="/coverage.html#portugal-spain@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/portugal-spain.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Russia</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/russia.locus.xml">Locus</a></td>
  <td>1.8 GB</td>
  <td><a href="/coverage.html#russia@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/russia.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Switzerland</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/switzerland.locus.xml">Locus</a></td>
  <td>269 MB</td>
  <td><a href="/coverage.html#switzerland@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/switzerland.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Ukraine</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/ukraine.locus.xml">Locus</a></td>
  <td>477 MB</td>
  <td><a href="/coverage.html#ukraine@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/ukraine.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Albania, Bosnia-Herzegovina, Bulgaria, Croatia, Hungary, Kosovo, Macedonia, Moldova, Montenegro, Romania, Serbia, Slovakia and Slovenia</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/al-ba-bg-hr-hu-xk-mk-md-me-ro-rs-sk-si.locus.xml">Locus</a></td>
  <td>1.0 GB</td>
  <td><a href="/coverage.html#al-ba-bg-hr-hu-xk-mk-md-me-ro-rs-sk-si@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/al-ba-bg-hr-hu-xk-mk-md-me-ro-rs-sk-si.osm-gh.zip">download</a>
</td>
</tr>
 
</table>
</div>

### Northern Europe

<div class="v-1.0-pre20">
  
<table>
<tr>
  <th>Region</th>
  <th>Install</th>
  <th>Size</th>
  <th>Coverage</th>
</tr>

<tr>
  <td class="regionInfo">Denmark</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/denmark.locus.xml">Locus</a></td>
  <td>210 MB</td>
  <td><a href="/coverage.html#denmark@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/denmark.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Finland</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/finland.locus.xml">Locus</a></td>
  <td>316 MB</td>
  <td><a href="/coverage.html#finland@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/finland.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Finland, Norway and Sweden</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/finland-norway-sweden.locus.xml">Locus</a></td>
  <td>943 MB</td>
  <td><a href="/coverage.html#finland-norway-sweden@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/finland-norway-sweden.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Great Britain</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/great-britain.locus.xml">Locus</a></td>
  <td>1.0 GB</td>
  <td><a href="/coverage.html#great-britain@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/great-britain.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Iceland</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/iceland.locus.xml">Locus</a></td>
  <td>21 MB</td>
  <td><a href="/coverage.html#iceland@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/iceland.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Ireland and Northern Ireland</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/ireland-and-northern-ireland.locus.xml">Locus</a></td>
  <td>145 MB</td>
  <td><a href="/coverage.html#ireland-and-northern-ireland@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/ireland-and-northern-ireland.osm-gh.zip">download</a>
</td>
</tr>
 
</table>
</div>

<div class="v-1.0-pre26" style="display: none">
  
<table>
<tr>
  <th>Region</th>
  <th>Install</th>
  <th>Size</th>
  <th>Coverage</th>
</tr>

<tr>
  <td class="regionInfo">Denmark</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/denmark.locus.xml">Locus</a></td>
  <td>211 MB</td>
  <td><a href="/coverage.html#denmark@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/denmark.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Finland</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/finland.locus.xml">Locus</a></td>
  <td>319 MB</td>
  <td><a href="/coverage.html#finland@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/finland.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Finland, Norway and Sweden</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/finland-norway-sweden.locus.xml">Locus</a></td>
  <td>953 MB</td>
  <td><a href="/coverage.html#finland-norway-sweden@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/finland-norway-sweden.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Great Britain</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/great-britain.locus.xml">Locus</a></td>
  <td>1.0 GB</td>
  <td><a href="/coverage.html#great-britain@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/great-britain.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Iceland</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/iceland.locus.xml">Locus</a></td>
  <td>21 MB</td>
  <td><a href="/coverage.html#iceland@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/iceland.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Ireland and Northern Ireland</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/ireland-and-northern-ireland.locus.xml">Locus</a></td>
  <td>144 MB</td>
  <td><a href="/coverage.html#ireland-and-northern-ireland@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/ireland-and-northern-ireland.osm-gh.zip">download</a>
</td>
</tr>
 
</table>
</div>

<div class="v-1.0-pre31" style="display: none">
  
<table>
<tr>
  <th>Region</th>
  <th>Install</th>
  <th>Size</th>
  <th>Coverage</th>
</tr>

<tr>
  <td class="regionInfo">Denmark</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/denmark.locus.xml">Locus</a></td>
  <td>210 MB</td>
  <td><a href="/coverage.html#denmark@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/denmark.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Finland</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/finland.locus.xml">Locus</a></td>
  <td>317 MB</td>
  <td><a href="/coverage.html#finland@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/finland.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Finland, Norway and Sweden</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/finland-norway-sweden.locus.xml">Locus</a></td>
  <td>949 MB</td>
  <td><a href="/coverage.html#finland-norway-sweden@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/finland-norway-sweden.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Great Britain</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/great-britain.locus.xml">Locus</a></td>
  <td>998 MB</td>
  <td><a href="/coverage.html#great-britain@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/great-britain.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Iceland</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/iceland.locus.xml">Locus</a></td>
  <td>21 MB</td>
  <td><a href="/coverage.html#iceland@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/iceland.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Ireland and Northern Ireland</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/ireland-and-northern-ireland.locus.xml">Locus</a></td>
  <td>142 MB</td>
  <td><a href="/coverage.html#ireland-and-northern-ireland@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/ireland-and-northern-ireland.osm-gh.zip">download</a>
</td>
</tr>
 
</table>
</div>

### North America

<div class="v-1.0-pre20">
  
<table>
<tr>
  <th>Region</th>
  <th>Install</th>
  <th>Size</th>
  <th>Coverage</th>
</tr>

<tr>
  <td class="regionInfo">Canada</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/canada.locus.xml">Locus</a></td>
  <td>630 MB</td>
  <td><a href="/coverage.html#canada@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/canada.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">US Midwest</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/us-midwest.locus.xml">Locus</a></td>
  <td>1.7 GB</td>
  <td><a href="/coverage.html#us-midwest@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/us-midwest.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">US Northeast</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/us-northeast.locus.xml">Locus</a></td>
  <td>807 MB</td>
  <td><a href="/coverage.html#us-northeast@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/us-northeast.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">US Pacific</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/us-pacific.locus.xml">Locus</a></td>
  <td>25 MB</td>
  <td><a href="/coverage.html#us-pacific@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/us-pacific.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">US South</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/us-south.locus.xml">Locus</a></td>
  <td>2.5 GB</td>
  <td><a href="/coverage.html#us-south@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip:
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/us-south-part1.osm-gh.zip">part 1</a> (591 MB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/us-south-part2.osm-gh.zip">part 2</a> (562 MB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/us-south-part3.osm-gh.zip">part 3</a> (1.4 GB)
    
</td>
</tr>

<tr>
  <td class="regionInfo">US West</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/us-west.locus.xml">Locus</a></td>
  <td>1.7 GB</td>
  <td><a href="/coverage.html#us-west@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/us-west.osm-gh.zip">download</a>
</td>
</tr>
 
</table>
</div>

<div class="v-1.0-pre26" style="display: none">
  
<table>
<tr>
  <th>Region</th>
  <th>Install</th>
  <th>Size</th>
  <th>Coverage</th>
</tr>

<tr>
  <td class="regionInfo">Canada</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/canada.locus.xml">Locus</a></td>
  <td>636 MB</td>
  <td><a href="/coverage.html#canada@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/canada.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">US Midwest</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/us-midwest.locus.xml">Locus</a></td>
  <td>1.7 GB</td>
  <td><a href="/coverage.html#us-midwest@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/us-midwest.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">US Northeast</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/us-northeast.locus.xml">Locus</a></td>
  <td>816 MB</td>
  <td><a href="/coverage.html#us-northeast@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/us-northeast.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">US Pacific</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/us-pacific.locus.xml">Locus</a></td>
  <td>25 MB</td>
  <td><a href="/coverage.html#us-pacific@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/us-pacific.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">US South</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/us-south.locus.xml">Locus</a></td>
  <td>2.5 GB</td>
  <td><a href="/coverage.html#us-south@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip:
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/us-south-part1.osm-gh.zip">part 1</a> (598 MB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/us-south-part2.osm-gh.zip">part 2</a> (567 MB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/us-south-part3.osm-gh.zip">part 3</a> (1.4 GB)
    
</td>
</tr>

<tr>
  <td class="regionInfo">US West</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/us-west.locus.xml">Locus</a></td>
  <td>1.7 GB</td>
  <td><a href="/coverage.html#us-west@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/us-west.osm-gh.zip">download</a>
</td>
</tr>
 
</table>
</div>

<div class="v-1.0-pre31" style="display: none">
  
<table>
<tr>
  <th>Region</th>
  <th>Install</th>
  <th>Size</th>
  <th>Coverage</th>
</tr>

<tr>
  <td class="regionInfo">Canada</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/canada.locus.xml">Locus</a></td>
  <td>633 MB</td>
  <td><a href="/coverage.html#canada@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/canada.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">US Midwest</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/us-midwest.locus.xml">Locus</a></td>
  <td>1.6 GB</td>
  <td><a href="/coverage.html#us-midwest@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/us-midwest.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">US Northeast</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/us-northeast.locus.xml">Locus</a></td>
  <td>803 MB</td>
  <td><a href="/coverage.html#us-northeast@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/us-northeast.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">US Pacific</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/us-pacific.locus.xml">Locus</a></td>
  <td>25 MB</td>
  <td><a href="/coverage.html#us-pacific@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/us-pacific.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">US South</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/us-south.locus.xml">Locus</a></td>
  <td>2.5 GB</td>
  <td><a href="/coverage.html#us-south@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip:
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/us-south-part1.osm-gh.zip">part 1</a> (999 MB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/us-south-part2.osm-gh.zip">part 2</a> (571 MB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/us-south-part3.osm-gh.zip">part 3</a> (906 MB)
    
</td>
</tr>

<tr>
  <td class="regionInfo">US West</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/us-west.locus.xml">Locus</a></td>
  <td>1.7 GB</td>
  <td><a href="/coverage.html#us-west@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/us-west.osm-gh.zip">download</a>
</td>
</tr>
 
</table>
</div>

### Asia

<div class="v-1.0-pre20">
  
<table>
<tr>
  <th>Region</th>
  <th>Install</th>
  <th>Size</th>
  <th>Coverage</th>
</tr>

<tr>
  <td class="regionInfo">China</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/china.locus.xml">Locus</a></td>
  <td>1.0 GB</td>
  <td><a href="/coverage.html#china@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/china.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">India</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/india.locus.xml">Locus</a></td>
  <td>979 MB</td>
  <td><a href="/coverage.html#india@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/india.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Indonesia</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/indonesia.locus.xml">Locus</a></td>
  <td>954 MB</td>
  <td><a href="/coverage.html#indonesia@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/indonesia.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Japan</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/japan.locus.xml">Locus</a></td>
  <td>2.0 GB</td>
  <td><a href="/coverage.html#japan@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip:
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/japan-part1.osm-gh.zip">part 1</a> (540 MB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/japan-part2.osm-gh.zip">part 2</a> (507 MB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/japan-part3.osm-gh.zip">part 3</a> (992 MB)
    
</td>
</tr>

<tr>
  <td class="regionInfo">Taiwan</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/taiwan.locus.xml">Locus</a></td>
  <td>135 MB</td>
  <td><a href="/coverage.html#taiwan@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/taiwan.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Thailand</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/thailand.locus.xml">Locus</a></td>
  <td>463 MB</td>
  <td><a href="/coverage.html#thailand@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/thailand.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Turkey</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/turkey.locus.xml">Locus</a></td>
  <td>564 MB</td>
  <td><a href="/coverage.html#turkey@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/turkey.osm-gh.zip">download</a>
</td>
</tr>
 
</table>
</div>

<div class="v-1.0-pre26" style="display: none">
  
<table>
<tr>
  <th>Region</th>
  <th>Install</th>
  <th>Size</th>
  <th>Coverage</th>
</tr>

<tr>
  <td class="regionInfo">China</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/china.locus.xml">Locus</a></td>
  <td>1.1 GB</td>
  <td><a href="/coverage.html#china@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/china.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">India</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/india.locus.xml">Locus</a></td>
  <td>988 MB</td>
  <td><a href="/coverage.html#india@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/india.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Indonesia</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/indonesia.locus.xml">Locus</a></td>
  <td>970 MB</td>
  <td><a href="/coverage.html#indonesia@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/indonesia.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Japan</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/japan.locus.xml">Locus</a></td>
  <td>2.1 GB</td>
  <td><a href="/coverage.html#japan@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip:
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/japan-part1.osm-gh.zip">part 1</a> (544 MB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/japan-part2.osm-gh.zip">part 2</a> (510 MB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/japan-part3.osm-gh.zip">part 3</a> (997 MB)
    
</td>
</tr>

<tr>
  <td class="regionInfo">Nepal</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/nepal.locus.xml">Locus</a></td>
  <td>124 MB</td>
  <td><a href="/coverage.html#nepal@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/nepal.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Philippines</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/philippines.locus.xml">Locus</a></td>
  <td>178 MB</td>
  <td><a href="/coverage.html#philippines@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/philippines.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Taiwan</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/taiwan.locus.xml">Locus</a></td>
  <td>136 MB</td>
  <td><a href="/coverage.html#taiwan@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/taiwan.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Thailand</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/thailand.locus.xml">Locus</a></td>
  <td>463 MB</td>
  <td><a href="/coverage.html#thailand@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/thailand.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Turkey</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/turkey.locus.xml">Locus</a></td>
  <td>568 MB</td>
  <td><a href="/coverage.html#turkey@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/turkey.osm-gh.zip">download</a>
</td>
</tr>
 
</table>
</div>

<div class="v-1.0-pre31" style="display: none">
  
<table>
<tr>
  <th>Region</th>
  <th>Install</th>
  <th>Size</th>
  <th>Coverage</th>
</tr>

<tr>
  <td class="regionInfo">China</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/china.locus.xml">Locus</a></td>
  <td>1.1 GB</td>
  <td><a href="/coverage.html#china@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/china.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">India</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/india.locus.xml">Locus</a></td>
  <td>997 MB</td>
  <td><a href="/coverage.html#india@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/india.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Indonesia</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/indonesia.locus.xml">Locus</a></td>
  <td>976 MB</td>
  <td><a href="/coverage.html#indonesia@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/indonesia.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Japan</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/japan.locus.xml">Locus</a></td>
  <td>2.0 GB</td>
  <td><a href="/coverage.html#japan@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip:
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/japan-part1.osm-gh.zip">part 1</a> (956 MB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/japan-part2.osm-gh.zip">part 2</a> (512 MB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/japan-part3.osm-gh.zip">part 3</a> (582 MB)
    
</td>
</tr>

<tr>
  <td class="regionInfo">Nepal</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/nepal.locus.xml">Locus</a></td>
  <td>127 MB</td>
  <td><a href="/coverage.html#nepal@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/nepal.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Philippines</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/philippines.locus.xml">Locus</a></td>
  <td>180 MB</td>
  <td><a href="/coverage.html#philippines@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/philippines.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Taiwan</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/taiwan.locus.xml">Locus</a></td>
  <td>138 MB</td>
  <td><a href="/coverage.html#taiwan@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/taiwan.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Thailand</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/thailand.locus.xml">Locus</a></td>
  <td>461 MB</td>
  <td><a href="/coverage.html#thailand@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/thailand.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Turkey</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/turkey.locus.xml">Locus</a></td>
  <td>570 MB</td>
  <td><a href="/coverage.html#turkey@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/turkey.osm-gh.zip">download</a>
</td>
</tr>
 
</table>
</div>

### Other

<div class="v-1.0-pre20">
  
<table>
<tr>
  <th>Region</th>
  <th>Install</th>
  <th>Size</th>
  <th>Coverage</th>
</tr>

<tr>
  <td class="regionInfo">Africa</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/africa.locus.xml">Locus</a></td>
  <td>4.4 GB</td>
  <td><a href="/coverage.html#africa@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip:
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/africa-part1.osm-gh.zip">part 1</a> (1.0 GB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/africa-part2.osm-gh.zip">part 2</a> (993 MB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/africa-part3.osm-gh.zip">part 3</a> (2.4 GB)
    
</td>
</tr>

<tr>
  <td class="regionInfo">Australia</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/australia.locus.xml">Locus</a></td>
  <td>541 MB</td>
  <td><a href="/coverage.html#australia@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/australia.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Brazil</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/brazil.locus.xml">Locus</a></td>
  <td>1.5 GB</td>
  <td><a href="/coverage.html#brazil@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/brazil.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Central America</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/central-america.locus.xml">Locus</a></td>
  <td>364 MB</td>
  <td><a href="/coverage.html#central-america@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/central-america.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Morocco</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/morocco.locus.xml">Locus</a></td>
  <td>270 MB</td>
  <td><a href="/coverage.html#morocco@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/morocco.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">New Zealand</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/new-zealand.locus.xml">Locus</a></td>
  <td>79 MB</td>
  <td><a href="/coverage.html#new-zealand@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/new-zealand.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">South America</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-02-03/south-america.locus.xml">Locus</a></td>
  <td>2.8 GB</td>
  <td><a href="/coverage.html#south-america@1.0-pre20">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip:
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/south-america-part1.osm-gh.zip">part 1</a> (744 MB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/south-america-part2.osm-gh.zip">part 2</a> (678 MB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-02-03/south-america-part3.osm-gh.zip">part 3</a> (1.4 GB)
    
</td>
</tr>
 
</table>
</div>

<div class="v-1.0-pre26" style="display: none">
  
<table>
<tr>
  <th>Region</th>
  <th>Install</th>
  <th>Size</th>
  <th>Coverage</th>
</tr>

<tr>
  <td class="regionInfo">Africa</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/africa.locus.xml">Locus</a></td>
  <td>4.5 GB</td>
  <td><a href="/coverage.html#africa@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip:
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/africa-part1.osm-gh.zip">part 1</a> (1.1 GB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/africa-part2.osm-gh.zip">part 2</a> (1.0 GB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/africa-part3.osm-gh.zip">part 3</a> (2.4 GB)
    
</td>
</tr>

<tr>
  <td class="regionInfo">Australia</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/australia.locus.xml">Locus</a></td>
  <td>544 MB</td>
  <td><a href="/coverage.html#australia@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/australia.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Brazil</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/brazil.locus.xml">Locus</a></td>
  <td>1.5 GB</td>
  <td><a href="/coverage.html#brazil@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/brazil.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Central America</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/central-america.locus.xml">Locus</a></td>
  <td>366 MB</td>
  <td><a href="/coverage.html#central-america@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/central-america.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Morocco</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/morocco.locus.xml">Locus</a></td>
  <td>270 MB</td>
  <td><a href="/coverage.html#morocco@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/morocco.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">New Zealand</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/new-zealand.locus.xml">Locus</a></td>
  <td>80 MB</td>
  <td><a href="/coverage.html#new-zealand@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/new-zealand.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">South America</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-02/south-america.locus.xml">Locus</a></td>
  <td>2.9 GB</td>
  <td><a href="/coverage.html#south-america@1.0-pre26">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip:
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/south-america-part1.osm-gh.zip">part 1</a> (751 MB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/south-america-part2.osm-gh.zip">part 2</a> (684 MB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-02/south-america-part3.osm-gh.zip">part 3</a> (1.4 GB)
    
</td>
</tr>
 
</table>
</div>

<div class="v-1.0-pre31" style="display: none">
  
<table>
<tr>
  <th>Region</th>
  <th>Install</th>
  <th>Size</th>
  <th>Coverage</th>
</tr>

<tr>
  <td class="regionInfo">Africa</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/africa.locus.xml">Locus</a></td>
  <td>4.5 GB</td>
  <td><a href="/coverage.html#africa@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip:
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/africa-part1.osm-gh.zip">part 1</a> (1.9 GB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/africa-part2.osm-gh.zip">part 2</a> (1.0 GB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/africa-part3.osm-gh.zip">part 3</a> (1.6 GB)
    
</td>
</tr>

<tr>
  <td class="regionInfo">Australia</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/australia.locus.xml">Locus</a></td>
  <td>542 MB</td>
  <td><a href="/coverage.html#australia@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/australia.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Brazil</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/brazil.locus.xml">Locus</a></td>
  <td>1.5 GB</td>
  <td><a href="/coverage.html#brazil@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/brazil.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Central America</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/central-america.locus.xml">Locus</a></td>
  <td>366 MB</td>
  <td><a href="/coverage.html#central-america@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/central-america.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">Morocco</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/morocco.locus.xml">Locus</a></td>
  <td>268 MB</td>
  <td><a href="/coverage.html#morocco@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/morocco.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">New Zealand</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/new-zealand.locus.xml">Locus</a></td>
  <td>80 MB</td>
  <td><a href="/coverage.html#new-zealand@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip: <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/new-zealand.osm-gh.zip">download</a>
</td>
</tr>

<tr>
  <td class="regionInfo">South America</td>
  <td><a href="locus-actions://https/gh-data.org/locus/2020-03-21/south-america.locus.xml">Locus</a></td>
  <td>2.9 GB</td>
  <td><a href="/coverage.html#south-america@1.0-pre31">coverage</a></td>
</tr>

<tr class="infoRow">
<td colSpan="4">
  zip:
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/south-america-part1.osm-gh.zip">part 1</a> (1.3 GB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/south-america-part2.osm-gh.zip">part 2</a> (688 MB),
    <a href="https://s3.eu-central-1.wasabisys.com/gh-routing-data/2020-03-21/south-america-part3.osm-gh.zip">part 3</a> (918 MB)
    
</td>
</tr>
 
</table>
</div>



<!-- end of generated block -->