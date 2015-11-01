var Root = React.createClass({
  getInitialState: function() {
    return {candy_names: []};
  },
  componentDidMount: function() {
    $.getJSON("/api/candies", function(data) {
      this.setState({candy_names: data["candies"]});
    }.bind(this));
  },
  render: function() {
    return (
      <div className="container candies">
        <div className="row">
          <CandiesList names={this.state.candy_names}/>
        </div>
        <div className="row">
          <h3>Create New Candy</h3>
          <CandyCreateForm />
        </div>
      </div>
    );
  }
});

var CandiesList = React.createClass({
  render: function() {
    var listElts = [];
    for(var i = 0; i < this.props.names.length; i++) {
      listElts.push(<li className="list-group-item">{this.props.names[i]}</li>);
    }
    return (
      <ul className="list-group">
        {listElts}
      </ul>
    );
  }
});

var CandyCreateForm = React.createClass({
  render: function() {
    return (
      <form>
        <div className="form-group">
          <label for="candyName">Candy Name</label>
          <input type="text" className="form-control" id="candyName" placeholder="Candy Name"/>
        </div>
        <button type="submit" className="btn btn-default">Create</button>
      </form>
    );
  }
})

ReactDOM.render(<Root />, document.getElementById("page-content"));
