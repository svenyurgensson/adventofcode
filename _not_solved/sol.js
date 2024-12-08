
const input = await Deno.readTextFile( 'input.txt' );

const by = { x: new Map, y: new Map };
const start = {};
const max = {};
// заполним справочные структуры данных
{
  const fill = ( [ m, m_key, val ] ) => {
    const before = m.get( m_key );
    if( before )
      before.f.push( val );
    else
      m.set( m_key, { f: [ val ] } );
  };
  for( const [ y, line ] of input.split( /\r?\n/ ).entries() ){
    if( line == "" ) break;
    for( const [ x, ch ] of line.split( '' ).entries() ){
      switch( ch ){
        case '#':
          [ [ by.x, x, y ]
          , [ by.y, y, x ] ].forEach( fill );
          break;
        case '^':
          start.x = x; start.y = y;
          break;
      }
      max.x = x;
    }
    max.y = y;
  }
  for( const i of [ 'x', 'y' ] ){
    for( const s of by[ i ].values() ){
      s.r = [ ...s.f ].reverse();
    }
  }
}
// console.dir( { by, start, max } );

const add_2dim_bitmap = ( mem, { x, y } ) => {
  const before = mem.get( x );
  if( before )
    before.add( y );
  else
    mem.set( x, new Set( [ y ] ) );
};
const get_2dim_bitmap = ( mem, { x, y } ) => {
  return mem.get( x ) && mem.get( x ).has( y );
};
const del_2dim_bitmap = ( mem, { x, y } ) => {
  mem.get( x ).delete( y );
};

const ffind = ( where, first ) => where.f.find( i => i > first );
const rfind = ( where, first ) => where.r.find( i => i < first );
// n - север, e - восток, s - юг, w - запад
const dir_n = { name: 'n', turns: new Map, main: 'x', walk: 'y', find: rfind, back:  1, border:     0 };
const dir_e = { name: 'e', turns: new Map, main: 'y', walk: 'x', find: ffind, back: -1, border: max.x };
const dir_s = { name: 's', turns: new Map, main: 'x', walk: 'y', find: ffind, back: -1, border: max.y };
const dir_w = { name: 'w', turns: new Map, main: 'y', walk: 'x', find: rfind, back:  1, border:     0 };
dir_n.next = dir_e; dir_e.next = dir_s; dir_s.next = dir_w; dir_w.next = dir_n;

const next_obstruction = ( { main, walk, find }, pos, trying ) => {
  const same_line_obstructions = by[ main ].get( pos[ main ] ) || { f: [], r: [] };
  const nearest = find( same_line_obstructions, pos[ walk ] );
  if( trying && trying[ main ] == pos[ main ] ){
    switch( true ){
      case nearest == null && find( { f: [ trying[ walk ] ], r: [ trying[ walk ] ] }, pos[ walk ] ) != null:
      case pos[ walk ] > trying[ walk ] && trying[ walk ] > nearest:
      case pos[ walk ] < trying[ walk ] && trying[ walk ] < nearest:
        return trying[ walk ];
    }
  }
  return nearest;
};

const checked = new Map;
let result = 0;
const leads_outside = ( dir, pos, trying = null ) => {
  const nearest = next_obstruction( dir, pos, trying );
  const { main, walk, back, turns } = dir;
  if( trying == null ){
    const check_last = ( nearest == null ) ? dir.border - back : nearest;
    for( let t = pos[walk] - back; t != check_last; t = t - back ){
      const checked_pos = { [ main ]: pos[ main ], [ walk ]: t };
      if( get_2dim_bitmap( checked, checked_pos ) ) continue;
      if( ! leads_outside( dir, pos, checked_pos ) ){ result++; }
      add_2dim_bitmap( checked, checked_pos );
    }
  }
  if( nearest == null ) return true;
  const next_pos = { [ main ]: pos[ main ], [ walk ]: nearest + back }; // мы в петле
  if( get_2dim_bitmap( turns, next_pos ) ) return false;
  add_2dim_bitmap( turns, next_pos );
  const deeper = leads_outside( dir.next, next_pos, trying );
  del_2dim_bitmap( turns, next_pos );
  return deeper;
};

leads_outside( dir_n, start );
console.dir( result );