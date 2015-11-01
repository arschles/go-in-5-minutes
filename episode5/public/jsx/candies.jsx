var CandiesList = React.createClass({
  getInitialState: function() {
    return {keys: []}
  },
  componentDidMount: function() {
    $.getJSON("/api/candy_keys", function(data) {
      this.setState({keys: data["keys"]});
    }.bind(this));
  },
  render: function() {
    return (
      <div className="container candies">
        <div className="row">
          <CandiesList keys={this.state.keys}/>
        </div>
        <div className="row">
          <h3>Create New Candy</h3>
          <CandyCreateForm />
        </div>
      </div>
    );
  }
});

var CandiesList = React.createClass({});

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

ReactDOM.render(<CandiesList />, document.getElementById("page-content"));
